package runner

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	coreumchainconfig "github.com/CoreumFoundation/coreum/v4/pkg/config"
	coreumchainconstant "github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

const (
	configVersion = "v1"
	// ConfigFileName is the file name used for the client config.
	ConfigFileName = "client.yaml"
	// DefaultCoreumChainID is default chain id.
	// TODO: Change to ChainIDMain before release
	DefaultCoreumChainID = coreumchainconstant.ChainIDTest
)

// Runner is client runner that aggregates all client components.
type Runner struct {
	cfg           Config
	log           logger.Logger
	components    Components
	clientAddress sdk.AccAddress
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

	return &Runner{
		cfg:           cfg,
		log:           components.Log,
		components:    components,
		clientAddress: clientAddress,
	}, nil
}

// Start starts runner.
func (r *Runner) Start(ctx context.Context) error {
	// TODO: Write main procedure here
	return nil
}

// Components groups components required by runner.
type Components struct {
	Log                  logger.Logger
	RunnerConfig         Config
	CoreumSDKClientCtx   client.Context
	CoreumClientCtx      coreumchainclient.Context
	CoreumContractClient *coreum.ContractClient
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

	return Components{
		Log:                  log,
		RunnerConfig:         cfg,
		CoreumSDKClientCtx:   coreumSDKClientCtx,
		CoreumClientCtx:      coreumClientCtx,
		CoreumContractClient: contractClient,
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
