package runner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/CoreumFoundation/iso20022-client/iso20022/runner"
)

//nolint:tparallel // the test is parallel, but test cases are not
func TestInitAndReadConfig(t *testing.T) {
	t.Parallel()

	defaultCfg := runner.DefaultConfig()
	yamlStringConfig, err := yaml.Marshal(defaultCfg)
	require.NoError(t, err)
	require.Equal(t, getDefaultConfigString(), string(yamlStringConfig))

	tests := []struct {
		name                  string
		beforeWriteModifyFunc func(config runner.Config) runner.Config
		expectedConfigFunc    func(config runner.Config) runner.Config
	}{
		{
			name:                  "default_config",
			beforeWriteModifyFunc: func(config runner.Config) runner.Config { return config },
			expectedConfigFunc:    func(config runner.Config) runner.Config { return config },
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			// not parallel intentionally top prevent race

			// create temp dir to store the config
			tempDir := tt.TempDir()
			// try to read none-existing config
			_, err = runner.ReadConfigFromFile(tempDir)
			require.Error(tt, err)

			// init the config first time
			modifiedCfg := tc.beforeWriteModifyFunc(defaultCfg)
			require.NoError(tt, runner.InitConfig(tempDir, modifiedCfg))

			// try to init the config second time
			require.Error(tt, runner.InitConfig(tempDir, modifiedCfg))

			// read config
			readConfig, err := runner.ReadConfigFromFile(tempDir)
			require.NoError(tt, err)
			require.Error(tt, runner.InitConfig(tempDir, defaultCfg))

			require.Equal(tt, tc.expectedConfigFunc(defaultCfg), readConfig)
		})
	}
}

// the func returns the default config snapshot as string.
func getDefaultConfigString() string {
	return `version: v1
logging:
    level: info
    format: console
coreum:
    client_key_name: iso20022-client
    grpc:
        url: https://full-node.testnet-1.coreum.dev:9090
    network:
        chain_id: coreum-testnet-1
    contract:
        contract_address: testcore1za96naulkx2axrq738x9uke65ztq2grffuyds67kzwms75tj8lfq9272g0
        gas_adjustment: 1.4
        gas_price_adjustment: 1.2
        page_limit: 50
        out_of_gas_retry_delay: 500ms
        out_of_gas_retry_attempts: 5
        request_timeout: 10s
        tx_timeout: 1m0s
        tx_status_poll_interval: 500ms
`
}
