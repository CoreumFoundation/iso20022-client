package processes_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	integrationtests "github.com/CoreumFoundation/iso20022-client/integration-tests"
	"github.com/CoreumFoundation/iso20022-client/iso20022/queue"
)

func TestProcesses(t *testing.T) {
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

	_, err := firstPartyRunnerEnv.SendMessage("../../iso20022/messages/testdata/pacs008-1.xml")
	requireT.NoError(err)
	<-time.After(15 * time.Second) // Wait a bit till the message is received
	msg, err := secondPartyRunnerEnv.ReceiveMessage()
	requireT.NoError(err)

	requireT.NotEmpty(msg)

	_, err = firstPartyRunnerEnv.MessageStatus("P5607186 298")
	requireT.ErrorContains(err, "message not found")

	status, err := firstPartyRunnerEnv.SendMessage("../../iso20022/messages/testdata/pacs008-2.xml")
	requireT.NoError(err)
	requireT.Equal(queue.StatusSending, status.DeliveryStatus)

	<-time.After(15 * time.Second) // Wait a bit till the message is received

	status, err = secondPartyRunnerEnv.MessageStatus("P5607186 298")
	requireT.NoError(err)
	requireT.Equal(queue.StatusSent, status.DeliveryStatus)

	msg, err = secondPartyRunnerEnv.ReceiveMessage()
	requireT.NoError(err)

	requireT.NotEmpty(msg)
}
