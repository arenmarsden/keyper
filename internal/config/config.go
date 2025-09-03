package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const ReadWriteOnly = 0o600

type Config struct {
	StorageProvider string `mapstructure:"storage_provider"`
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Region          string `mapstructure:"region"`
	UseSSL          bool   `mapstructure:"use_ssl"`
}

func getConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".config/keyper")
}

func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigPermissions(ReadWriteOnly)
	v.AddConfigPath("$HOME/.config/keyper")
	v.AddConfigPath(".")
	return v
}

func LoadConfig() (*Config, error) {
	v := InitViper()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := os.MkdirAll(getConfigDir(), 0o755); err != nil {
				return nil, err
			}
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func WriteConfig(cfg *Config) error {
	v := InitViper()

	v.Set("storage_provider", cfg.StorageProvider)
	v.Set("endpoint", cfg.Endpoint)
	v.Set("access_key_id", cfg.AccessKeyID)
	v.Set("secret_access_key", cfg.SecretAccessKey)
	v.Set("region", cfg.Region)
	v.Set("use_ssl", cfg.UseSSL)

	return v.WriteConfigAs(filepath.Join(os.Getenv("HOME"), ".config/keyper/config.yaml"))
}
