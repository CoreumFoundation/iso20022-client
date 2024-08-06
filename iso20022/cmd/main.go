package main

import (
	"context"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/CoreumFoundation/coreum-tools/pkg/run"
	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	"github.com/CoreumFoundation/coreum/v4/pkg/config"
	"github.com/CoreumFoundation/iso20022-client/iso20022/cmd/cli"
)

// @title ISO20022 Client
// @version 0.0.1

// @BasePath /v1
func main() {
	run.Tool("iso20022", func(ctx context.Context) error {
		rootCmd, err := RootCmd(ctx)
		if err != nil {
			return err
		}

		err = rootCmd.Execute()
		if err != nil && !errors.Is(err, context.Canceled) {
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

	cmd.AddCommand(cli.InitCmd())
	cmd.AddCommand(cli.StartCmd(processorProvider))
	cmd.AddCommand(cli.VersionCmd())
	cmd.AddCommand(cli.MessageCmd())

	keyringCmd, err := cli.KeyringCmd()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(keyringCmd)

	return cmd, nil
}

func processorProvider(cmd *cobra.Command) (cli.Runner, error) {
	rnr, err := cli.NewRunnerFromHome(cmd)
	if err != nil {
		return nil, err
	}

	return rnr, nil
}
