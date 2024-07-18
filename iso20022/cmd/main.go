package main

import (
	"context"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/CoreumFoundation/coreum-tools/pkg/run"
	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	"github.com/CoreumFoundation/coreum/v4/pkg/config"
	coreumchainconstant "github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
	"github.com/CoreumFoundation/iso20022-client/iso20022/cmd/cli"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
)

func main() {
	run.Tool("iso20022", func(ctx context.Context) error {
		rootCmd, err := RootCmd(ctx)
		if err != nil {
			return err
		}
		if err := rootCmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
			return err
		}

		return nil
	})
}

// RootCmd returns the root cmd.
//
//nolint:contextcheck // the context is passed in the command
func RootCmd(ctx context.Context) (*cobra.Command, error) {
	encodingConfig := config.NewEncodingConfig(coreumapp.ModuleBasics)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := &cobra.Command{
		Short: "Iso20022 client.",
	}
	cmd.SetContext(ctx)
	cli.AddHomeFlag(cmd)
	cmd.InitDefaultHelpFlag()

	if cfg, err := cli.GetHomeRunnerConfig(cmd); err == nil && cfg.Coreum.Network.ChainID != "" {
		coreumChainNetworkConfig, err := config.NetworkConfigByChainID(
			coreumchainconstant.ChainID(cfg.Coreum.Network.ChainID),
		)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get coreum network config for chainID:%s", cfg.Coreum.Network.ChainID)
		}
		coreum.SetSDKConfig(coreumChainNetworkConfig.Provider.GetAddressPrefix())
	} else {
		return nil, errors.New("failed to get chainID from config")
	}

	cmd.AddCommand(cli.InitCmd())
	cmd.AddCommand(cli.StartCmd(processorProvider))
	cmd.AddCommand(keys.Commands(cli.DefaultHomeDir))
	cmd.AddCommand(cli.VersionCmd())
	cmd.AddCommand(cli.MessageCmd())

	return cmd, nil
}

func processorProvider(cmd *cobra.Command) (cli.Runner, error) {
	rnr, err := cli.NewRunnerFromHome(cmd)
	if err != nil {
		return nil, err
	}

	return rnr, nil
}
