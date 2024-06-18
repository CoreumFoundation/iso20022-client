//nolint:tagliatelle // yaml naming
package runner

import (
	"io"
	"os"
	"path/filepath"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

// LoggingConfig is logging config.
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// HTTPClientConfig is http client config.
type HTTPClientConfig struct {
	RequestTimeout time.Duration `yaml:"request_timeout"`
	DoTimeout      time.Duration `yaml:"do_timeout"`
	RetryDelay     time.Duration `yaml:"retry_delay"`
}

// CoreumGRPCConfig is the coreum GRPC config.
type CoreumGRPCConfig struct {
	URL string `yaml:"url"`
}

// CoreumNetworkConfig is coreum network config.
type CoreumNetworkConfig struct {
	ChainID string `yaml:"chain_id"`
}

// CoreumContractConfig is coreum contract config.
type CoreumContractConfig struct {
	ContractAddress       string        `yaml:"contract_address"`
	GasAdjustment         float64       `yaml:"gas_adjustment"`
	GasPriceAdjustment    float64       `yaml:"gas_price_adjustment"`
	PageLimit             uint32        `yaml:"page_limit"`
	OutOfGasRetryDelay    time.Duration `yaml:"out_of_gas_retry_delay"`
	OutOfGasRetryAttempts uint32        `yaml:"out_of_gas_retry_attempts"`
	// client context config
	RequestTimeout       time.Duration `yaml:"request_timeout"`
	TxTimeout            time.Duration `yaml:"tx_timeout"`
	TxStatusPollInterval time.Duration `yaml:"tx_status_poll_interval"`
}

// CoreumConfig is coreum config.
type CoreumConfig struct {
	ClientKeyName string               `yaml:"client_key_name"`
	GRPC          CoreumGRPCConfig     `yaml:"grpc"`
	Network       CoreumNetworkConfig  `yaml:"network"`
	Contract      CoreumContractConfig `yaml:"contract"`
}

// ProcessesConfig  is processes config.
type ProcessesConfig struct {
	RepeatDelay time.Duration `yaml:"repeat_delay"`
	RetryDelay  time.Duration `yaml:"retry_delay"`
	ExitOnError bool          `yaml:"-"`
}

// MetricsPeriodicCollectorConfig is metric periodic collector config.
type MetricsPeriodicCollectorConfig struct {
	RepeatDelay time.Duration `yaml:"repeat_delay"`
}

// Config is runner config.
type Config struct {
	Version       string        `yaml:"version"`
	LoggingConfig LoggingConfig `yaml:"logging"`
	Coreum        CoreumConfig  `yaml:"coreum"`
}

// DefaultConfig returns default runner config.
func DefaultConfig() Config {
	defaultCoreumContactConfig := coreum.DefaultContractClientConfig(sdk.AccAddress{})
	defaultClientCtxDefaultCfg := coreumchainclient.DefaultContextConfig()

	defaultLoggerConfig := logger.DefaultZapLoggerConfig()

	return Config{
		Version: configVersion,
		LoggingConfig: LoggingConfig{
			Level:  defaultLoggerConfig.Level,
			Format: defaultLoggerConfig.Format,
		},

		Coreum: CoreumConfig{
			ClientKeyName: "iso20022-client",
			GRPC: CoreumGRPCConfig{
				// TODO: Change to mainnet url before release
				URL: "https://full-node.testnet-1.coreum.dev:9090",
			},
			Network: CoreumNetworkConfig{
				ChainID: string(DefaultCoreumChainID),
			},
			Contract: CoreumContractConfig{
				// TODO: Change to the contract address on mainnet before release
				ContractAddress:       "testcore1za96naulkx2axrq738x9uke65ztq2grffuyds67kzwms75tj8lfq9272g0",
				GasAdjustment:         defaultCoreumContactConfig.GasAdjustment,
				GasPriceAdjustment:    defaultCoreumContactConfig.GasPriceAdjustment.MustFloat64(),
				PageLimit:             defaultCoreumContactConfig.PageLimit,
				OutOfGasRetryDelay:    defaultCoreumContactConfig.OutOfGasRetryDelay,
				OutOfGasRetryAttempts: defaultCoreumContactConfig.OutOfGasRetryAttempts,

				RequestTimeout:       defaultClientCtxDefaultCfg.TimeoutConfig.RequestTimeout,
				TxTimeout:            defaultClientCtxDefaultCfg.TimeoutConfig.TxTimeout,
				TxStatusPollInterval: defaultClientCtxDefaultCfg.TimeoutConfig.TxStatusPollInterval,
			},
		},
	}
}

// InitConfig creates config yaml file.
func InitConfig(homePath string, cfg Config) error {
	path := BuildFilePath(homePath)
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return errors.Errorf("failed to init config, file already exists, path:%s", path)
	}

	err := os.MkdirAll(homePath, 0o700)
	if err != nil {
		return errors.Errorf("failed to create dirs by path:%s", path)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return errors.Wrapf(err, "failed to create config file, path:%s", path)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	yamlStringConfig, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "failed convert default config to yaml")
	}
	if _, err := file.Write(yamlStringConfig); err != nil {
		return errors.Wrapf(err, "failed to write yaml config file, path:%s", path)
	}

	return nil
}

// ReadConfigFromFile reads config yaml file.
func ReadConfigFromFile(homePath string) (Config, error) {
	path := BuildFilePath(homePath)
	file, err := os.OpenFile(path, os.O_RDONLY, 0o600)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.Errorf("config file does not exist, path:%s", path)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, errors.Wrapf(err, "failed to read bytes from file does not exist, path:%s", path)
	}

	var config Config
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		return Config{}, errors.Wrapf(err, "failed to unmarshal file to yaml, path:%s", path)
	}

	return config, nil
}

// BuildFilePath builds the file path.
func BuildFilePath(homePath string) string {
	return filepath.Join(homePath, ConfigFileName)
}
