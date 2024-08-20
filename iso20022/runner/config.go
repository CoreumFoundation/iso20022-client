//nolint:tagliatelle // yaml naming
package runner

import (
	"io"
	"os"
	"path"
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
	Denom   string `yaml:"denom"`
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

// ServerConfig is the server config.
type ServerConfig struct {
	ListenAddress string `yaml:"listen_address"`
}

// AddressBookConfig is the address book config.
type AddressBookConfig struct {
	UpdateInterval    time.Duration `yaml:"update_interval"`
	CustomRepoAddress string        `yaml:"custom_repo_address"`
}

// DtifConfig is the dtif config.
type DtifConfig struct {
	UpdateInterval      time.Duration `yaml:"update_interval"`
	CustomSourceAddress string        `yaml:"custom_source_address"`
}

// Queue is the message queue config.
type Queue struct {
	Size                int           `yaml:"size"`
	Path                string        `yaml:"path"`
	StatusCacheDuration time.Duration `yaml:"status_cache_duration"`
}

// ProcessesConfig is processes config.
type ProcessesConfig struct {
	Server      ServerConfig      `yaml:"server"`
	AddressBook AddressBookConfig `yaml:"address_book"`
	Dtif        DtifConfig        `yaml:"dtif"`
	Queue       Queue             `yaml:"queue"`
	RetryDelay  time.Duration     `yaml:"retry_delay"`
}

// Config is runner config.
type Config struct {
	Version       string          `yaml:"version"`
	LoggingConfig LoggingConfig   `yaml:"logging"`
	Coreum        CoreumConfig    `yaml:"coreum"`
	Processes     ProcessesConfig `yaml:"processes"`
}

// DefaultConfig returns default runner config.
func DefaultConfig() Config {
	defaultCoreumContactConfig := coreum.DefaultContractClientConfig(sdk.AccAddress{})
	defaultClientCtxDefaultCfg := coreumchainclient.DefaultContextConfig()

	defaultLoggerConfig := logger.DefaultZapLoggerConfig()

	config := Config{
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
				Denom:   DefaultDenom,
			},
			Contract: CoreumContractConfig{
				// TODO: Change to the contract address on mainnet before release
				ContractAddress:       "testcore1434davf92sz8fpgc0rp3pyc8umxqk685he47xmz4ph7ujqz5vyys2frmdq",
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

		Processes: ProcessesConfig{
			Server: ServerConfig{
				ListenAddress: ":2843",
			},
			AddressBook: AddressBookConfig{
				UpdateInterval: 60 * time.Second,
			},
			Dtif: DtifConfig{
				UpdateInterval: 60 * time.Second,
			},
			Queue: Queue{
				Size:                50,
				Path:                path.Join(os.TempDir(), "iso20022"),
				StatusCacheDuration: time.Hour,
			},
			RetryDelay: 10 * time.Second,
		},
	}
	config.Processes.Queue.Path = path.Join(os.TempDir(), config.Coreum.ClientKeyName)
	return config
}

// InitConfig creates config yaml file.
func InitConfig(homePath string, cfg Config) error {
	filePath := BuildFilePath(homePath)
	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		return errors.Errorf("failed to init config, file already exists, path:%s", filePath)
	}

	err := os.MkdirAll(homePath, 0o700)
	if err != nil {
		return errors.Errorf("failed to create dirs by path:%s", filePath)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return errors.Wrapf(err, "failed to create config file, path:%s", filePath)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	yamlStringConfig, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "failed convert default config to yaml")
	}
	if _, err := file.Write(yamlStringConfig); err != nil {
		return errors.Wrapf(err, "failed to write yaml config file, path:%s", filePath)
	}

	return nil
}

// ReadConfigFromFile reads config yaml file.
func ReadConfigFromFile(homePath string) (Config, error) {
	filePath := BuildFilePath(homePath)
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o600)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.Errorf("config file does not exist, path:%s", filePath)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, errors.Wrapf(err, "failed to read bytes from file does not exist, path:%s", filePath)
	}

	var config Config
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		return Config{}, errors.Wrapf(err, "failed to unmarshal file to yaml, path:%s", filePath)
	}

	return config, nil
}

// BuildFilePath builds the file path.
func BuildFilePath(homePath string) string {
	return filepath.Join(homePath, ConfigFileName)
}
