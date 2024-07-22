package processes_test

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/coreum-tools/pkg/parallel"
	coreumintegration "github.com/CoreumFoundation/coreum/v4/testutil/integration"
	integrationtests "github.com/CoreumFoundation/iso20022-client/integration-tests"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/runner"
)

// RunnerEnvConfig is runner environment config.
type RunnerEnvConfig struct {
	AwaitTimeout          time.Duration
	CustomContractAddress *sdk.AccAddress
	CustomContractOwner   *sdk.AccAddress
	AccountMnemonics      string
}

// DefaultRunnerEnvConfig returns default runner environment config.
func DefaultRunnerEnvConfig() RunnerEnvConfig {
	return RunnerEnvConfig{
		AwaitTimeout:          5 * time.Minute,
		CustomContractAddress: nil,
		CustomContractOwner:   nil,
		AccountMnemonics:      "",
	}
}

// RunnerEnv is runner environment used for the integration tests.
type RunnerEnv struct {
	Cfg                  RunnerEnvConfig
	ContractClient       *coreum.ContractClient
	Chain                integrationtests.Chain
	ContractOwner        sdk.AccAddress
	RunnersParallelGroup *parallel.Group
	Runner               *runner.Runner
	RunnerComponent      runner.Components
	AccountAddress       sdk.AccAddress
}

// NewRunnerEnv returns new instance of the RunnerEnv.
func NewRunnerEnv(ctx context.Context, t *testing.T, cfg RunnerEnvConfig, chain integrationtests.Chain) *RunnerEnv {
	ctx, cancel := context.WithCancel(ctx)

	var contractOwner sdk.AccAddress
	if cfg.CustomContractOwner == nil {
		contractOwner = chain.Coreum.GenAccount()
	} else {
		contractOwner = *cfg.CustomContractOwner
	}

	// fund to cover the fees
	chain.Coreum.FundAccountWithOptions(ctx, t, contractOwner, coreumintegration.BalancesOptions{
		Amount: sdkmath.NewIntFromUint64(100_000_000),
	})

	contractClient := coreum.NewContractClient(
		coreum.DefaultContractClientConfig(sdk.AccAddress(nil)),
		chain.Log,
		chain.Coreum.ClientContext,
	)

	if cfg.CustomContractAddress == nil {
		contractAddress, err := contractClient.DeployAndInstantiate(ctx, contractOwner, chain.Coreum.Config().ContractPath)
		require.NoError(t, err)
		require.NoError(t, contractClient.SetContractAddress(contractAddress))
	} else {
		require.NoError(t, contractClient.SetContractAddress(*cfg.CustomContractAddress))
	}

	rnrComponents, rnr, accountAddress := createDevRunner(
		ctx,
		t,
		chain,
		contractClient.GetContractAddress(),
		cfg.AccountMnemonics,
	)

	runnerEnv := &RunnerEnv{
		Cfg:                  cfg,
		ContractClient:       contractClient,
		Chain:                chain,
		ContractOwner:        contractOwner,
		RunnersParallelGroup: parallel.NewGroup(ctx),
		Runner:               rnr,
		RunnerComponent:      rnrComponents,
		AccountAddress:       accountAddress,
	}
	t.Cleanup(func() {
		// we can cancel the context now and wait for the runner to stop gracefully
		cancel()
		err := runnerEnv.RunnersParallelGroup.Wait()
		if err == nil || errors.Is(err, context.Canceled) {
			return
		}
		// the client replies with that error if the context is canceled at the time of the request,
		// and the error is in the internal package, so we can't check the type
		if strings.Contains(err.Error(), "context canceled") {
			return
		}

		require.NoError(t, err, "Found unexpected runner process errors after the execution")
	})

	return runnerEnv
}

// StartRunnerProcesses starts all processes.
func (r *RunnerEnv) StartRunnerProcesses() {
	r.RunnersParallelGroup.Spawn("runner", parallel.Exit, r.Runner.Start)
}

func fundAccount(
	ctx context.Context,
	t *testing.T,
	coreumChain integrationtests.CoreumChain,
	keyName, mnemonic string,
) sdk.AccAddress {
	t.Helper()

	keyInfo, err := coreumChain.ClientContext.Keyring().NewAccount(
		keyName,
		mnemonic,
		"",
		hd.CreateHDPath(coreumChain.ChainSettings.CoinType, 0, 0).String(),
		hd.Secp256k1,
	)
	if err != nil {
		panic(err)
	}

	address, err := keyInfo.GetAddress()
	if err != nil {
		panic(err)
	}

	//address := coreumChain.ImportMnemonic(mnemonic)
	coreumChain.Faucet.FundAccounts(ctx, t, coreumintegration.FundedAccount{
		Address: address,
		Amount:  coreumChain.NewCoin(sdkmath.NewIntFromUint64(1_000_000)),
	})

	return address
}

func createDevRunner(
	ctx context.Context,
	t *testing.T,
	chain integrationtests.Chain,
	contractAddress sdk.AccAddress,
	accountMnemonics string,
) (runner.Components, *runner.Runner, sdk.AccAddress) {
	t.Helper()

	uniqueName := uniqueNameFromMnemonic(accountMnemonics)
	accountAddress := fundAccount(ctx, t, chain.Coreum, uniqueName, accountMnemonics)
	chain.Log.Info(ctx, "Account imported", zap.String("address", accountAddress.String()))

	runnerCfg := runner.DefaultConfig()
	runnerCfg.LoggingConfig.Level = "info"
	runnerCfg.Coreum.ClientKeyName = uniqueName
	runnerCfg.Coreum.GRPC.URL = chain.Coreum.Config().GRPCAddress
	runnerCfg.Coreum.Contract.ContractAddress = contractAddress.String()
	runnerCfg.Coreum.Network.ChainID = chain.Coreum.ChainSettings.ChainID
	runnerCfg.Processes.RepeatDelay = 500 * time.Millisecond
	runnerCfg.Processes.AddressBook.CustomRepoAddress = chain.Coreum.Config().AddressBookRepoAddress
	port, err := getFreePort()
	require.NoError(t, err)
	runnerCfg.Processes.Queue.Path = path.Join(os.TempDir(), uniqueName)
	runnerCfg.Processes.Server.ListenAddress = ":" + strconv.Itoa(port)

	// exit on errors
	runnerCfg.Processes.ExitOnError = true

	// re-init log to use correct `CallerSkip`
	log, err := logger.NewZapLogger(logger.DefaultZapLoggerConfig())
	require.NoError(t, err)

	coreumSDKClientCtx := chain.Coreum.ClientContext.SDKContext()
	components, err := runner.NewComponents(runnerCfg, coreumSDKClientCtx, log)
	require.NoError(t, err)

	rnr, err := runner.NewRunner(components, runnerCfg)
	require.NoError(t, err)
	return components, rnr, accountAddress
}

func uniqueNameFromMnemonic(mnemonic string) string {
	return fmt.Sprintf("iso20022-integration-test-%x", md5.Sum([]byte(mnemonic)))
}

func (r *RunnerEnv) SendMessage(messageFilePath string) error {
	file, err := os.OpenFile(messageFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// listen address here is always in the form of ":port", so we can append it to 0.0.0.0
	// to have full url of the listening service
	// the reason it is not just called port is that the RunnerConfig is the same between
	// integration-test and app
	uri := "http://0.0.0.0" + r.RunnerComponent.RunnerConfig.Processes.Server.ListenAddress
	res, err := http.Post(uri+"/v1/send", "application/xml", file)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusCreated {
		return errors.Errorf("http status %d: %s", res.StatusCode, res.Status)
	}

	return nil
}

func (r *RunnerEnv) ReceiveMessage() ([]byte, error) {
	// listen address here is always in the form of ":port", so we can append it to 0.0.0.0
	// to have full url of the listening service
	// the reason it is not just called port is that the RunnerConfig is the same between
	// integration-test and app
	uri := "http://0.0.0.0" + r.RunnerComponent.RunnerConfig.Processes.Server.ListenAddress
	res, err := http.Get(uri + "/v1/receive")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("http status %d: %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer func(l *net.TCPListener) {
				_ = l.Close()
			}(l)
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}
