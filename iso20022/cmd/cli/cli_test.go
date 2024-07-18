package cli

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	"github.com/CoreumFoundation/coreum/v4/pkg/config"
	"github.com/CoreumFoundation/iso20022-client/iso20022/runner"
)

func TestInitCmd(t *testing.T) {
	initConfig(t)
}

func TestStartCmd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	processorMock := NewMockRunner(ctrl)
	processorMock.EXPECT().Start(gomock.Any())
	cmd := StartCmd(func(cmd *cobra.Command) (Runner, error) {
		return processorMock, nil
	})
	executeCmd(t, cmd, initConfig(t)...)
}

func executeCmd(t *testing.T, cmd *cobra.Command, args ...string) string {
	return executeCmdWithOutputOption(t, cmd, "text", args...)
}

func executeCmdWithOutputOption(t *testing.T, cmd *cobra.Command, outOpt string, args ...string) string {
	t.Helper()

	cmd.SetArgs(args)

	buf := new(bytes.Buffer)
	cmd.SetErr(buf)
	cmd.SetOut(buf)
	cmd.SetArgs(args)

	encodingConfig := config.NewEncodingConfig(coreumapp.ModuleBasics)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithOutputFormat(outOpt)
	ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

	if err := cmd.ExecuteContext(ctx); err != nil {
		require.NoError(t, err)
	}

	t.Logf("Command %s is executed successfully", cmd.Name())

	return buf.String()
}

func flagWithPrefix(f string) string {
	return fmt.Sprintf("--%s", f)
}

func initConfig(t *testing.T) []string {
	configPath := path.Join(t.TempDir(), "config-path")
	configFilePath := path.Join(configPath, runner.ConfigFileName)
	require.NoFileExists(t, configFilePath)

	args := []string{
		flagWithPrefix(FlagHome), configPath,
	}
	executeCmd(t, InitCmd(), args...)
	require.FileExists(t, configFilePath)

	return args
}
