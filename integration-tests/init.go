package integrationtests

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	coreumlogger "github.com/CoreumFoundation/coreum-tools/pkg/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

var chain Chain

// flag variables.
var (
	coreumCfg CoreumChainConfig
)

// Chain struct holds chain required for the testing.
type Chain struct {
	Coreum CoreumChain
	Log    logger.Logger
}

//nolint:lll // breaking down cli flags will make it less readable.
func init() {
	flag.StringVar(&coreumCfg.GRPCAddress, "coreum-grpc-address", "localhost:9090", "GRPC address of cored node started by coreum")
	flag.StringVar(&coreumCfg.FundingMnemonic, "coreum-funding-mnemonic", "sad hobby filter tray ordinary gap half web cat hard call mystery describe member round trend friend beyond such clap frozen segment fan mistake", "Funding coreum account mnemonic required by tests")
	flag.StringVar(&coreumCfg.ContractPath, "coreum-contract-path", "../contract/iso_messaging_poc.wasm", "Path to smart contract wasm file")
	flag.StringVar(&coreumCfg.AddressBookRepoAddress, "address-book-repo-address", "file://../addressbook/addressbook.json", "Path to addressbook json file")
	flag.StringVar(&coreumCfg.Account1Mnemonic, "account1-mnemonic", "question minimum around dry mad beef vessel blouse submit lion woman twelve liquid enjoy replace river emerge process velvet stove hood tree minimum gun", "First account mnemonic")
	flag.StringVar(&coreumCfg.Account2Mnemonic, "account2-mnemonic", "genre plate metal lazy state govern panel scare clever broom yellow insane run easy turkey wool liberty core fire liquid menu cram toss outdoor", "Second account mnemonic")

	// accept testing flags
	testing.Init()
	// parse additional flags
	flag.Parse()

	logCfg := logger.DefaultZapLoggerConfig()
	// set the correct skip caller since we don't use the err counter wrapper here
	logCfg.CallerSkip = 1
	log, err := logger.NewZapLogger(logCfg)
	if err != nil {
		panic(errors.WithStack(err))
	}
	chain.Log = log

	coreumChain, err := NewCoreumChain(coreumCfg)
	if err != nil {
		panic(errors.Wrapf(err, "failed to init coreum chain"))
	}
	chain.Coreum = coreumChain

	// It just prevents some unwanted outputs while running the tests
	gin.SetMode(gin.ReleaseMode)
}

func mnemonicToTempPath(mnemonic string) string {
	return path.Join(os.TempDir(), fmt.Sprintf("iso20022-integration-test-%x", md5.Sum([]byte(mnemonic))))
}

// NewTestingContext returns the configured coreum chain and new context for the integration tests.
func NewTestingContext(t *testing.T) (context.Context, Chain) {
	ctx := coreumlogger.WithLogger(context.Background(), coreumlogger.New(coreumlogger.ToolDefaultConfig))
	testCtx, testCtxCancel := context.WithTimeout(ctx, 30*time.Minute)
	t.Cleanup(func() {
		require.NoError(t, testCtx.Err())
		testCtxCancel()
		err := os.RemoveAll(mnemonicToTempPath(chain.Coreum.Config().Account1Mnemonic))
		require.NoError(t, err)
		err = os.RemoveAll(mnemonicToTempPath(chain.Coreum.Config().Account2Mnemonic))
		require.NoError(t, err)
	})

	return testCtx, chain
}
