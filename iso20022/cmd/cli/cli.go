package cli

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo"
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
	// FlagChainID is chain-id flag.
	FlagChainID = "chain-id"
	// FlagCoreumGRPCURL is Coreum GRPC URL flag.
	FlagCoreumGRPCURL = "coreum-grpc-url"
	// FlagCoreumContractAddress is the address of the bridge smart contract.
	FlagCoreumContractAddress = "coreum-contract-address"
	// FlagServerAddress is the address that http server listens to flag.
	FlagServerAddress = "server-addr"
	// FlagCachePath is the path to save caches
	FlagCachePath = "cache-path"
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

	err = components.AddressBook.Validate()
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

	cachePath, err := cmd.Flags().GetString(FlagCachePath)
	if err == nil && cachePath != "" {
		cfg.Processes.Queue.Path = cachePath
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

			chainID, err := cmd.Flags().GetString(FlagChainID)
			if err != nil {
				return errors.Wrapf(err, "failed to read %s", FlagChainID)
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
	return cmd.PersistentFlags().String(FlagChainID, string(runner.DefaultCoreumChainID), "Coreum chain ID")
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

			err = providedRunner.Start(cmd.Context())
			if err != nil && errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		},
	}
	AddHomeFlag(cmd)
	AddKeyringFlags(cmd)
	AddKeyNameFlag(cmd)
	AddServerAddressFlag(cmd)
	AddCachePathFlag(cmd)

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

// AddCachePathFlag adds the cache-path flag to the command.
func AddCachePathFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagCachePath, "", "Cache path")
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
	// TODO: Should validate fully before sending and there should be a local batch in a WAL style
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
