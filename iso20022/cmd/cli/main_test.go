package cli

import (
	"testing"

	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/runner"
)

func TestMain(m *testing.M) {
	coreum.SetSDKConfig(string(runner.DefaultCoreumChainID))
}
