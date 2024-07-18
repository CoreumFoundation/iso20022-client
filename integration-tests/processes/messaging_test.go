package processes_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

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

	requireT.NoError(firstPartyRunnerEnv.SendMessage("../../iso20022/messages/testdata/pacs008-1.xml"))
	time.Sleep(10 * time.Second) // Wait a bit till the message is received
	msg, err := secondPartyRunnerEnv.ReceiveMessage()
	requireT.NoError(err)

	requireT.NotEmpty(msg)
}
