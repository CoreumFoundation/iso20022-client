package processes_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	coreumintegration "github.com/CoreumFoundation/coreum/v4/testutil/integration"
	integrationtests "github.com/CoreumFoundation/iso20022-client/integration-tests"
)

func TestMessaging(t *testing.T) {
	t.Parallel()
	requireT := require.New(t)

	ctx, chain := integrationtests.NewTestingContext(t)

	firstPartyRunnerEnvCfg := DefaultRunnerEnvConfig()
	firstPartyRunnerEnvCfg.AccountMnemonics = chain.Coreum.Config().Account1Mnemonic
	firstPartyRunnerEnv := NewRunnerEnv(ctx, t, firstPartyRunnerEnvCfg, chain)
	firstPartyRunnerEnv.StartRunnerProcesses()

	secondPartyRunnerEnvCfg := DefaultRunnerEnvConfig()
	secondPartyRunnerEnvCfg.AccountMnemonics = chain.Coreum.Config().Account2Mnemonic
	secondPartyRunnerEnvCfg.CustomContractAddress = lo.ToPtr(firstPartyRunnerEnv.ContractClient.GetContractAddress())
	secondPartyRunnerEnvCfg.CustomContractOwner = lo.ToPtr(firstPartyRunnerEnv.ContractOwner)
	secondPartyRunnerEnv := NewRunnerEnv(ctx, t, secondPartyRunnerEnvCfg, chain)
	secondPartyRunnerEnv.StartRunnerProcesses()

	coreumSenderAddress := chain.Coreum.GenAccount()
	issueFee := chain.Coreum.QueryAssetFTParams(ctx, t).IssueFee
	chain.Coreum.FundAccountWithOptions(ctx, t, coreumSenderAddress, coreumintegration.BalancesOptions{
		Amount: issueFee.Amount.Add(sdkmath.NewIntWithDecimal(1, 6)),
	})

	requireT.NoError(firstPartyRunnerEnv.SendMessage("../../iso20022/messages/testdata/pacs008-1.xml"))
	msg, err := secondPartyRunnerEnv.ReceiveMessage(time.Minute)
	requireT.NoError(err)

	requireT.NotEmpty(msg)
}
