package cli

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo"
	overridekeyring "github.com/CoreumFoundation/iso20022-client/iso20022/cmd/cli/cosmos/override/crypto/keyring"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/runner"
)

//go:generate mockgen -destination=cli_mocks_test.go -package=cli . Runner

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultHomeDir = filepath.Join(userHomeDir, ".iso20022-client")
}

// DefaultHomeDir is default home for the iso client.
var DefaultHomeDir string

const (
	// CoreumKeyringSuffix is the Coreum keyring suffix.
	CoreumKeyringSuffix = "coreum"
)

const (
	// FlagHome is home flag.
	FlagHome = "home"
	// FlagKeyName is the key name flag.
	FlagKeyName = "key-name"
	// FlagCoreumChainID is chain-id flag.
	FlagCoreumChainID = "coreum-chain-id"
	// FlagCoreumGRPCURL is Coreum GRPC URL flag.
	FlagCoreumGRPCURL = "coreum-grpc-url"
	// FlagCoreumContractAddress is the address of the bridge smart contract.
	FlagCoreumContractAddress = "coreum-contract-address"
	// FlagServerAddress is the address that http server listens to flag.
	FlagServerAddress = "server-addr"
)

// Runner is a runner interface.
type Runner interface {
	Start(ctx context.Context) error
}

// RunnerProvider is a function that returns the Runner from the input cmd.
type RunnerProvider func(cmd *cobra.Command) (Runner, error)

// NewRunnerFromHome returns runner from home.
func NewRunnerFromHome(cmd *cobra.Command) (*runner.Runner, error) {
	cfg, err := GetHomeRunnerConfig(cmd)
	if err != nil {
		return nil, err
	}

	logCfg := logger.DefaultZapLoggerConfig()
	logCfg.Level = cfg.LoggingConfig.Level
	logCfg.Format = cfg.LoggingConfig.Format
	zapLogger, err := logger.NewZapLogger(logCfg)
	if err != nil {
		return nil, err
	}

	components, err := NewComponents(cmd, zapLogger)
	if err != nil {
		return nil, err
	}

	err = components.AddressBook.Update(cmd.Context())
	if err != nil {
		return nil, err
	}

	rnr, err := runner.NewRunner(components, cfg)
	if err != nil {
		return nil, err
	}

	return rnr, nil
}

// NewComponents creates components based on CLI input.
func NewComponents(cmd *cobra.Command, log logger.Logger) (runner.Components, error) {
	cfg, err := GetHomeRunnerConfig(cmd)
	if err != nil {
		return runner.Components{}, err
	}

	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return runner.Components{}, errors.Wrap(err, "failed to get client context")
	}
	coreumClientCtx, err := withKeyring(clientCtx, cmd.Flags(), CoreumKeyringSuffix, log)
	if err != nil {
		return runner.Components{}, errors.Wrap(err, "failed to configure coreum keyring")
	}

	components, err := runner.NewComponents(cfg, coreumClientCtx, log)
	if err != nil {
		return runner.Components{}, err
	}

	return components, nil
}

// withKeyring adds suffix-specific keyring witch decoded private key caching to the context.
func withKeyring(
	clientCtx client.Context,
	flagSet *pflag.FlagSet,
	suffix string,
	log logger.Logger,
) (client.Context, error) {
	if flagSet.Lookup(flags.FlagKeyringDir) == nil || flagSet.Lookup(flags.FlagKeyringBackend) == nil {
		return clientCtx, nil
	}
	keyringDir, err := flagSet.GetString(flags.FlagKeyringDir)
	if err != nil {
		return client.Context{}, errors.WithStack(err)
	}
	if keyringDir == "" {
		keyringDir = filepath.Join(clientCtx.HomeDir, "keyring")
	}
	keyringDir += "-" + suffix
	clientCtx = clientCtx.WithKeyringDir(keyringDir)

	keyringBackend, err := flagSet.GetString(flags.FlagKeyringBackend)
	if err != nil {
		return client.Context{}, errors.WithStack(err)
	}
	kr, err := client.NewKeyringFromBackend(clientCtx, keyringBackend)
	if err != nil {
		return client.Context{}, errors.WithStack(err)
	}

	return clientCtx.WithKeyring(newCacheKeyring(suffix, kr, clientCtx.Codec, log)), nil
}

// GetHomeRunnerConfig reads runner config from home directory.
func GetHomeRunnerConfig(cmd *cobra.Command) (runner.Config, error) {
	home, err := getClientHome(cmd)
	if err != nil {
		return runner.Config{}, err
	}

	cfg, err := runner.ReadConfigFromFile(home)
	if err != nil {
		return runner.Config{}, err
	}

	keyName, err := cmd.Flags().GetString(FlagKeyName)
	if err == nil && keyName != "" {
		cfg.Coreum.ClientKeyName = keyName
	}

	listenAddr, err := cmd.Flags().GetString(FlagServerAddress)
	if err == nil && listenAddr != "" {
		cfg.Processes.Server.ListenAddress = listenAddr
	}

	return cfg, nil
}

func getClientHome(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(FlagHome)
}

// GetCLILogger returns the console logger initialized with the default logger config but with set `yaml` format.
func GetCLILogger() (*logger.ZapLogger, error) {
	zapLogger, err := logger.NewZapLogger(logger.ZapLoggerConfig{
		Level:  "info",
		Format: logger.YamlConsoleLoggerFormat,
	})
	if err != nil {
		return nil, err
	}

	return zapLogger, nil
}

// InitCmd returns the init cmd.
func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the iso 20022 client home with the default config",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			home, err := getClientHome(cmd)
			if err != nil {
				return err
			}
			log, err := GetCLILogger()
			if err != nil {
				return err
			}
			log.Info(ctx, "Generating settings", zap.String("home", home))

			chainID, err := cmd.Flags().GetString(FlagCoreumChainID)
			if err != nil {
				return errors.Wrapf(err, "failed to read %s", FlagCoreumChainID)
			}
			coreumGRPCURL, err := cmd.Flags().GetString(FlagCoreumGRPCURL)
			if err != nil {
				return errors.Wrapf(err, "failed to read %s", FlagCoreumGRPCURL)
			}
			coreumContractAddress, err := cmd.Flags().GetString(FlagCoreumContractAddress)
			if err != nil {
				return errors.Wrapf(err, "failed to read %s", FlagCoreumContractAddress)
			}

			cfg := runner.DefaultConfig()
			cfg.Coreum.Network.ChainID = chainID
			cfg.Coreum.GRPC.URL = coreumGRPCURL
			cfg.Coreum.Contract.ContractAddress = coreumContractAddress

			keyName, err := cmd.Flags().GetString(FlagKeyName)
			if err == nil && keyName != "" {
				cfg.Coreum.ClientKeyName = keyName
			}

			if err = runner.InitConfig(home, cfg); err != nil {
				return err
			}
			log.Info(ctx, "Settings are generated successfully")
			return nil
		},
	}

	addCoreumChainIDFlag(cmd)
	cmd.PersistentFlags().String(FlagCoreumGRPCURL, "", "Coreum GRPC address.")
	cmd.PersistentFlags().String(FlagCoreumContractAddress, "", "Address of the smart contract.")

	AddHomeFlag(cmd)

	return cmd
}

func addCoreumChainIDFlag(cmd *cobra.Command) *string {
	return cmd.PersistentFlags().String(FlagCoreumChainID, string(runner.DefaultCoreumChainID), "Default coreum chain ID")
}

// StartCmd returns the start cmd.
func StartCmd(pp RunnerProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the client",
		RunE: func(cmd *cobra.Command, args []string) error {
			providedRunner, err := pp(cmd)
			if err != nil {
				return err
			}

			return providedRunner.Start(cmd.Context())
		},
	}
	AddHomeFlag(cmd)
	AddKeyringFlags(cmd)
	AddKeyNameFlag(cmd)
	AddServerAddressFlag(cmd)

	return cmd
}

// AddHomeFlag adds the home flag to the command.
func AddHomeFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagHome, DefaultHomeDir, "Client home directory")
}

// AddKeyringFlags adds keyring flags to the command.
func AddKeyringFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(
		flags.FlagKeyringBackend,
		flags.DefaultKeyringBackend,
		"Select keyring backend (os|file|kwallet|pass|test)",
	)
	cmd.PersistentFlags().String(
		flags.FlagKeyringDir,
		"", "The client Keyring directory; if omitted, the default 'home' directory will be used")
}

// AddKeyNameFlag adds the key-name flag to the command.
func AddKeyNameFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagKeyName, "", "Key name from the keyring")
}

// AddServerAddressFlag adds the server-addr flag to the command.
func AddServerAddressFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagServerAddress, "", "Http server address")
}

// ClientKeysCmd prints the client keys info.
func ClientKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client-keys",
		Short: "Print the Coreum keys info",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			log, err := GetCLILogger()
			if err != nil {
				return err
			}

			components, err := NewComponents(cmd, log)
			if err != nil {
				return err
			}

			// Coreum
			coreumKeyRecord, err := components.CoreumClientCtx.Keyring().Key(components.RunnerConfig.Coreum.ClientKeyName)
			if err != nil {
				return errors.Wrapf(err, "failed to get coreum key, keyName:%s", components.RunnerConfig.Coreum.ClientKeyName)
			}
			coreumAddress, err := coreumKeyRecord.GetAddress()
			if err != nil {
				return errors.Wrapf(err, "failed to get coreum address from key, keyName:%s",
					components.RunnerConfig.Coreum.ClientKeyName)
			}

			pubKey := "unknown"

			switch coreumKeyRecord.PubKey.TypeUrl {
			case "/cosmos.crypto.secp256k1.PubKey":
				var key secp256k1.PubKey
				err = proto.Unmarshal(coreumKeyRecord.PubKey.Value, &key)
				if err != nil {
					return err
				}
				pubKey = base64.StdEncoding.EncodeToString(key.Key)
			case "/cosmos.crypto.secp256r1.PubKey":
				var key secp256r1.PubKey
				err = proto.Unmarshal(coreumKeyRecord.PubKey.Value, &key)
				if err != nil {
					return err
				}
				pubKeyBytes, err := x509.MarshalPKIXPublicKey(key.Key.PublicKey)
				if err != nil {
					return err
				}
				pubKey = base64.StdEncoding.EncodeToString(pubKeyBytes)
			}

			components.Log.Info(
				ctx,
				"Keys info",
				zap.String("coreumAddress", coreumAddress.String()),
				zap.String("publicKey", pubKey),
			)

			return nil
		},
	}
	AddKeyringFlags(cmd)
	AddKeyNameFlag(cmd)
	AddHomeFlag(cmd)

	return cmd
}

// KeyringCmd returns cosmos keyring cmd init with the correct keys home.
func KeyringCmd(
	coinType uint32,
	addressFormatter overridekeyring.AddressFormatter,
) (*cobra.Command, error) {
	// We need to set CoinType before initializing keys commands because keys.Commands() sets default
	// flag value from sdk config. See github.com/cosmos/cosmos-sdk@v0.47.5/client/keys/add.go:78
	sdk.GetConfig().SetCoinType(coinType)

	// we set it for the keyring manually since it doesn't use the runner which does it for other CLI commands
	cmd := keys.Commands(DefaultHomeDir)
	for _, childCmd := range cmd.Commands() {
		childCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
			overridekeyring.SelectedAddressFormatter = addressFormatter

			log, err := GetCLILogger()
			if err != nil {
				return err
			}

			components, err := NewComponents(cmd, log)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContext(cmd, components.CoreumSDKClientCtx); err != nil {
				return errors.WithStack(err)
			}
			return nil
		}
	}

	return cmd, nil
}

// VersionCmd returns a CLI command to interactively print the application binary version information.
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the application binary version information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			log, err := GetCLILogger()
			if err != nil {
				return err
			}
			log.Info(
				cmd.Context(),
				"Version Info",
				zap.String("Git Tag", buildinfo.VersionTag),
				zap.String("Git Commit", buildinfo.GitCommit),
				zap.String("Build Time", buildinfo.BuildTime),
			)
			return nil
		},
	}
}

// MessageCmd returns the message cmd.
func MessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message",
		Short: "Send or receive ISO20022 messages",
	}
	AddHomeFlag(cmd)
	AddKeyringFlags(cmd)
	AddKeyNameFlag(cmd)
	AddServerAddressFlag(cmd)

	cmd.AddCommand(SendMessageCmd())
	cmd.AddCommand(ReceiveMessageCmd())

	return cmd
}

// SendMessageCmd returns a CLI command to interactively send an ISO20022 message to the chain.
func SendMessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send <message xml file>",
		Short: "Send an ISO20022 message to the chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log, err := GetCLILogger()
			if err != nil {
				return err
			}
			filePath := args[0]

			cfg, err := GetHomeRunnerConfig(cmd)
			if err != nil {
				return err
			}

			file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
			if err != nil {
				return err
			}

			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			res, err := http.Post(parseListenAddress(cfg.Processes.Server.ListenAddress)+"/v1/send", "application/xml", file)
			if err != nil {
				return err
			}

			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(res.Body)

			if res.StatusCode != http.StatusCreated {
				return errors.Errorf("http status %d: %s", res.StatusCode, res.Status)
			}

			log.Info(cmd.Context(), "Message sent")
			return nil
		},
	}

	AddServerAddressFlag(cmd)

	return cmd
}

// ReceiveMessageCmd returns a CLI command to interactively receive an ISO20022 message from the chain.
func ReceiveMessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive <path to save xml message>",
		Short: "Receive an ISO20022 message from the chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log, err := GetCLILogger()
			if err != nil {
				return err
			}
			filePath := args[0]

			cfg, err := GetHomeRunnerConfig(cmd)
			if err != nil {
				return err
			}

			file, err := os.Create(filePath)
			if err != nil {
				return err
			}

			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			res, err := http.Get(parseListenAddress(cfg.Processes.Server.ListenAddress) + "/v1/receive")
			if err != nil {
				return err
			}

			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(res.Body)

			if res.StatusCode == http.StatusNoContent {
				log.Info(cmd.Context(), "No new message")
				return nil
			}

			if res.StatusCode != http.StatusOK {
				return errors.Errorf("http status %d: %s", res.StatusCode, res.Status)
			}

			_, err = io.Copy(file, res.Body)
			if err != nil {
				return err
			}

			log.Info(cmd.Context(), "Message received")
			return nil
		},
	}

	AddServerAddressFlag(cmd)

	return cmd
}

func parseListenAddress(address string) string {
	defaultCfg := runner.DefaultConfig()
	parts := strings.Split(address, ":")
	host := "0.0.0.0"
	port := strings.Split(defaultCfg.Processes.Server.ListenAddress, ":")[1]
	if len(parts) > 0 && len(parts[0]) > 0 {
		host = parts[0]
	}
	if len(parts) > 1 && len(parts[1]) > 0 {
		port = parts[1]
	}
	return fmt.Sprintf("http://%s:%s", host, port)
}
