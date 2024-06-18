package cli

import (
	"bufio"
	"context"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/iso20022-client/iso20022/buildinfo"
	overridekeyring "github.com/CoreumFoundation/iso20022-client/iso20022/cmd/cli/cosmos/override/crypto/keyring"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

//go:generate mockgen -destination=cli_mocks_test.go -package=cli_test . Runner

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
	// FlagCoreumGRPCURL is Coreum GRPC URL flag.
	FlagCoreumGRPCURL = "coreum-grpc-url"
	// FlagCoreumContractAddress is the address of the bridge smart contract.
	FlagCoreumContractAddress = "coreum-contract-address"
	// FlagInitOnly is init only flag.
	FlagInitOnly = "init-only"
)

// Runner is a runner interface.
type Runner interface {
	Start(ctx context.Context) error
}

// RunnerProvider is a function that returns the Runner from the input cmd.
type RunnerProvider func(cmd *cobra.Command) (Runner, error)

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
			return nil
		},
	}

	cmd.PersistentFlags().String(FlagCoreumGRPCURL, "", "Coreum GRPC address.")
	cmd.PersistentFlags().String(FlagCoreumContractAddress, "", "Address of the smart contract.")

	AddHomeFlag(cmd)

	return cmd
}

// StartCmd returns the start cmd.
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the client",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			// Scan helps to wait for any input infinitely and just then call the client. That handles
			// the client restart in the container. Because after the restart, the container is detached, the client
			// requests the keyring password and fails immediately.
			log, err := GetCLILogger()
			if err != nil {
				return err
			}
			log.Info(ctx, "Press any key to start the client.")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			// TODO: Implement start here

			return nil
		},
	}
	AddHomeFlag(cmd)
	AddKeyringFlags(cmd)

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

// KeyringCmd returns cosmos keyring cmd inti with the correct keys home.
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
