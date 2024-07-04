package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	krflags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	coreumapp "github.com/CoreumFoundation/coreum/v4/app"
	"github.com/CoreumFoundation/coreum/v4/pkg/config"
	"github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
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

func TestKeyringCmds(t *testing.T) {
	cmd, err := KeyringCmd(constant.CoinType)
	require.NoError(t, err)

	args := append(initConfig(t), "list")
	args = append(args, testKeyringFlags(t.TempDir())...)

	out := executeCmd(t, cmd, args...)
	keysOut := make([]string, 0)
	require.NoError(t, json.Unmarshal([]byte(out), &keysOut))
	require.Empty(t, keysOut)
}

func TestKeyInfoCmd(t *testing.T) {
	keyringDir := t.TempDir()
	args := append(initConfig(t), testKeyringFlags(keyringDir)...)
	runnerDefaultCfg := runner.DefaultConfig()

	// add required key
	addKeyToTestKeyring(t, keyringDir, runnerDefaultCfg.Coreum.ClientKeyName, CoreumKeyringSuffix,
		sdk.GetConfig().GetFullBIP44Path())

	executeCmd(t, ClientKeysCmd(), args...)
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

func addKeyToTestKeyring(t *testing.T, keyringDir, keyName, suffix, hdPath string) sdk.AccAddress {
	keyringDir += "-" + suffix
	encodingConfig := config.NewEncodingConfig(coreumapp.ModuleBasics)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithOutputFormat("text").
		WithKeyringDir(keyringDir)

	kr, err := client.NewKeyringFromBackend(clientCtx, keyring.BackendTest)
	require.NoError(t, err)

	keyInfo, _, err := kr.NewMnemonic(
		keyName,
		keyring.English,
		hdPath,
		"",
		hd.Secp256k1,
	)
	require.NoError(t, err)

	addr, err := keyInfo.GetAddress()
	require.NoError(t, err)

	return addr
}

func testKeyringFlags(keyringDir string) []string {
	return []string{
		flagWithPrefix(krflags.FlagKeyringBackend), "test",
		flagWithPrefix(krflags.FlagKeyringDir), keyringDir,
	}
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
