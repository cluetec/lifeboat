package config

import (
	"fmt"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"strings"
	"zf.com/dos/backup/internal/destination/azureblob"
)

type Config struct {
	LogLevel    string
	Source      SourceConfig      `validate:"required"`
	Destination DestinationConfig `validate:"required"`
}

type SourceConfig struct {
	Type string `validate:"required"`
}

type HashiCorpVaultConfig struct {
	SourceConfig
	Address string
	Token   string
}

type DestinationConfig struct {
	AzureBlob azureblob.AzureBlobConfig
	Path      string
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func setDefaults(*viper.Viper) {
	viper.SetDefault("logLevel", "info")
}

func NewConfig() (*Config, error) {
	var config Config
	if err := defaults.Set(&config); err != nil {
		return &Config{}, fmt.Errorf("error while setting the default values to the config: %w", err)
	}

	v := viper.New()
	setDefaults(v)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath(".")
	v.SetConfigFile("./config.yaml")

	if err := v.ReadInConfig(); err != nil {
		return &Config{}, fmt.Errorf("error while reading in the configs: %w", err)
	}

	if err := v.Unmarshal(&config); err != nil {
		return &Config{}, fmt.Errorf("error while unmarshaling in the configs: %w", err)
	}

	validate = validator.New()
	if err := validate.Struct(config); err != nil {
		return &Config{}, fmt.Errorf("error while validating the configs: %w", err)
	}

	return &config, nil
}

func (c *Config) GetLogLevel() zerolog.Level {
	var level zerolog.Level

	switch strings.ToLower(c.LogLevel) {
	case "trace":
		level = zerolog.TraceLevel
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	return level
}
