package runner

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"runtime/debug"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/CoreumFoundation/coreum-tools/pkg/parallel"
	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	coreumchainconfig "github.com/CoreumFoundation/coreum/v4/pkg/config"
	coreumchainconstant "github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/compress"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/crypto"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/messages"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
	"github.com/CoreumFoundation/iso20022-client/iso20022/server"
)

const (
	configVersion = "v1"
	// ConfigFileName is the file name used for the client config.
	ConfigFileName = "client.yaml"
	// DefaultCoreumChainID is default chain id.
	// TODO: Change to ChainIDMain before release
	DefaultCoreumChainID = coreumchainconstant.ChainIDDev
)

// Runner is client runner that aggregates all client components.
type Runner struct {
	cfg           Config
	log           logger.Logger
	components    Components
	clientAddress sdk.AccAddress

	contractClientProcess     *processes.ContractClientProcess
	addressBookUpdaterProcess *processes.AddressBookUpdaterProcess
	webServer                 *server.Server
}

// NewRunner return new runner from the config.
func NewRunner(components Components, cfg Config) (*Runner, error) {
	if cfg.Coreum.Contract.ContractAddress == "" {
		return nil, errors.New("contract address is not configured")
	}

	clientAddress, err := getAddressFromKeyring(components.CoreumClientCtx.Keyring(), cfg.Coreum.ClientKeyName)
	if err != nil {
		return nil, err
	}

	addressBookUpdaterProcess, err := processes.NewAddressBookUpdaterProcess(
		cfg.Processes.AddressBook.UpdateInterval,
		components.Log,
		components.AddressBook,
	)
	if err != nil {
		return nil, err
	}

	sendCh := make(chan []byte, cfg.Processes.QueueSize)
	receiveCh := make(chan []byte, cfg.Processes.QueueSize)
	receiverProcess, err := processes.NewContractClientProcess(
		processes.ContractClientProcessConfig{
			CoreumContractAddress: components.CoreumContractClient.GetContractAddress(),
			ClientAddress:         clientAddress,
			ClientKeyName:         components.RunnerConfig.Coreum.ClientKeyName,
			PollInterval:          cfg.Processes.RetryDelay,
		},
		components.Log,
		components.Compressor,
		components.CoreumClientCtx,
		components.AddressBook,
		components.CoreumContractClient,
		components.Cryptography,
		components.Parser,
		sendCh,
		receiveCh,
	)
	if err != nil {
		return nil, err
	}

	handler := server.CreateHandlers(components.Parser, sendCh, receiveCh)
	webServer := server.New(cfg.Processes.Server.ListenAddress, handler)

	return &Runner{
		cfg:           cfg,
		log:           components.Log,
		components:    components,
		clientAddress: clientAddress,

		contractClientProcess:     receiverProcess,
		addressBookUpdaterProcess: addressBookUpdaterProcess,
		webServer:                 webServer,
	}, nil
}

// Start starts runner.
func (r *Runner) Start(ctx context.Context) error {
	runnerProcesses := map[string]func(context.Context) error{
		"contractClient": taskWithRestartOnError(
			r.contractClientProcess.Start,
			r.log,
			r.cfg.Processes.ExitOnError,
			r.cfg.Processes.RetryDelay,
		),
		"updateAddressBook": taskWithRestartOnError(
			r.addressBookUpdaterProcess.Start,
			r.log,
			r.cfg.Processes.ExitOnError,
			r.cfg.Processes.RetryDelay,
		),
		"webServer": taskWithRestartOnError(
			r.webServer.Start,
			r.log,
			true,
			r.cfg.Processes.RetryDelay,
		),
	}

	err := r.components.AddressBook.Update(ctx)
	if err != nil {
		return err
	}

	err = r.components.AddressBook.Validate()
	if err != nil {
		return err
	}

	return parallel.Run(ctx, func(ctx context.Context, spawn parallel.SpawnFn) error {
		for name, start := range runnerProcesses {
			name := name
			start := start
			spawn(name, parallel.Continue, start)
		}
		return nil
	})
}

func taskWithRestartOnError(
	task parallel.Task,
	log logger.Logger,
	exitOnError bool,
	retryDelay time.Duration,
) parallel.Task {
	return func(ctx context.Context) error {
		for {
			// start the process and handle the panic
			err := func() (err error) {
				defer func() {
					if p := recover(); p != nil {
						err = errors.Wrap(parallel.ErrPanic{Value: p, Stack: debug.Stack()}, "handled panic")
					}
				}()
				return task(ctx)
			}()

			if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil
			}

			// restart the process if it is restartable
			log.Error(ctx, "Received unexpected error from the process", zap.Error(err))
			if exitOnError {
				log.Warn(ctx, "The process is not auto-restartable on error")
				return err
			}

			if retryDelay > 0 {
				log.Info(ctx,
					"Process is paused and will be restarted later",
					zap.Duration("retryDelay", retryDelay))
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(retryDelay):
				}
			}
			log.Info(ctx, "Restarting the process after the error")
		}
	}
}

// Components groups components required by runner.
type Components struct {
	Log                  logger.Logger
	Compressor           *compress.Compressor
	RunnerConfig         Config
	CoreumSDKClientCtx   client.Context
	CoreumClientCtx      coreumchainclient.Context
	CoreumContractClient *coreum.ContractClient
	AddressBook          *addressbook.AddressBook
	Cryptography         *crypto.Cryptography
	Parser               *messages.Parser
}

// NewComponents creates components required by runner and other CLI commands.
func NewComponents(
	cfg Config,
	coreumSDKClientCtx client.Context,
	log logger.Logger,
) (Components, error) {
	coreumClientContextCfg := coreumchainclient.DefaultContextConfig()
	coreumClientContextCfg.TimeoutConfig.RequestTimeout = cfg.Coreum.Contract.RequestTimeout
	coreumClientContextCfg.TimeoutConfig.TxTimeout = cfg.Coreum.Contract.TxTimeout
	coreumClientContextCfg.TimeoutConfig.TxStatusPollInterval = cfg.Coreum.Contract.TxStatusPollInterval

	coreumClientCtx := coreumchainclient.NewContext(coreumClientContextCfg, coreumapp.ModuleBasics).
		WithKeyring(coreumSDKClientCtx.Keyring).
		WithGenerateOnly(coreumSDKClientCtx.GenerateOnly).
		WithFromAddress(coreumSDKClientCtx.FromAddress)

	if cfg.Coreum.Network.ChainID != "" {
		coreumChainNetworkConfig, err := coreumchainconfig.NetworkConfigByChainID(
			coreumchainconstant.ChainID(cfg.Coreum.Network.ChainID),
		)
		if err != nil {
			return Components{}, errors.Wrapf(
				err,
				"failed to set get correum network config for the chainID, chainID:%s",
				cfg.Coreum.Network.ChainID,
			)
		}
		coreumClientCtx = coreumClientCtx.WithChainID(cfg.Coreum.Network.ChainID)

		coreum.SetSDKConfig(coreumChainNetworkConfig.Provider.GetAddressPrefix())
	}

	var contractAddress sdk.AccAddress
	if cfg.Coreum.Contract.ContractAddress != "" {
		var err error
		contractAddress, err = sdk.AccAddressFromBech32(cfg.Coreum.Contract.ContractAddress)
		if err != nil {
			return Components{}, errors.Wrapf(
				err,
				"failed to decode contract address to sdk.AccAddress, address:%s",
				cfg.Coreum.Contract.ContractAddress,
			)
		}
	}
	contractClientCfg := coreum.DefaultContractClientConfig(contractAddress)
	contractClientCfg.GasAdjustment = cfg.Coreum.Contract.GasAdjustment
	contractClientCfg.GasPriceAdjustment = sdk.MustNewDecFromStr(fmt.Sprintf("%f", cfg.Coreum.Contract.GasPriceAdjustment))
	contractClientCfg.PageLimit = cfg.Coreum.Contract.PageLimit
	contractClientCfg.OutOfGasRetryDelay = cfg.Coreum.Contract.OutOfGasRetryDelay
	contractClientCfg.OutOfGasRetryAttempts = cfg.Coreum.Contract.OutOfGasRetryAttempts

	if cfg.Coreum.GRPC.URL != "" {
		grpcClient, err := getGRPCClientConn(cfg.Coreum.GRPC.URL)
		if err != nil {
			return Components{}, errors.Wrapf(err, "failed to create coreum GRPC client, URL:%s", cfg.Coreum.GRPC.URL)
		}
		coreumClientCtx = coreumClientCtx.WithGRPCClient(grpcClient)
	}

	contractClient := coreum.NewContractClient(contractClientCfg, log, coreumClientCtx)

	var addressBook *addressbook.AddressBook
	if cfg.Processes.AddressBook.CustomRepoAddress != "" {
		addressBook = addressbook.New(cfg.Coreum.Network.ChainID)
	} else {
		addressBook = addressbook.NewWithRepoAddress(cfg.Processes.AddressBook.CustomRepoAddress)
	}

	compressor, err := compress.New()
	if err != nil {
		return Components{}, err
	}

	cryptography := &crypto.Cryptography{}

	parser := messages.NewParser(log)

	return Components{
		Log:                  log,
		Compressor:           compressor,
		RunnerConfig:         cfg,
		CoreumSDKClientCtx:   coreumSDKClientCtx,
		CoreumClientCtx:      coreumClientCtx,
		CoreumContractClient: contractClient,
		AddressBook:          addressBook,
		Cryptography:         cryptography,
		Parser:               parser,
	}, nil
}

func getAddressFromKeyring(kr keyring.Keyring, keyName string) (sdk.AccAddress, error) {
	keyRecord, err := kr.Key(keyName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get key from the keyring, key name:%s", keyName)
	}
	addr, err := keyRecord.GetAddress()
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to get address from keyring key recodr, key name:%s",
			keyName,
		)
	}
	return addr, nil
}

func getGRPCClientConn(grpcURL string) (*grpc.ClientConn, error) {
	parsedURL, err := url.Parse(grpcURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse grpc URL")
	}

	encodingConfig := coreumchainconfig.NewEncodingConfig(coreumapp.ModuleBasics)
	pc, ok := encodingConfig.Codec.(codec.GRPCCodecProvider)
	if !ok {
		return nil, errors.New("failed to cast codec to codec.GRPCCodecProvider")
	}

	host := parsedURL.Host

	// https - tls grpc
	if parsedURL.Scheme == "https" {
		grpcClient, err := grpc.Dial(
			host,
			grpc.WithDefaultCallOptions(grpc.ForceCodec(pc.GRPCCodec())),
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to dial grpc")
		}
		return grpcClient, nil
	}

	// handling of host:port URL without the protocol
	if host == "" {
		host = fmt.Sprintf("%s:%s", parsedURL.Scheme, parsedURL.Opaque)
	}
	// http - insecure
	grpcClient, err := grpc.Dial(
		host,
		grpc.WithDefaultCallOptions(grpc.ForceCodec(pc.GRPCCodec())),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial grpc")
	}

	return grpcClient, nil
}
