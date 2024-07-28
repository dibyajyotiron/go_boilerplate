package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Consul   ConsulConfig
}

type DatabaseConfig struct {
	DSN string
}

type ConsulConfig struct {
	Address     string
	ServiceName string
	ServiceID   string
}

func LoadConfig() (*Config, error) {
	// Set environment to read from (default to "local")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
		os.Setenv("APP_ENV", env)
	}
	// Load environment variables from .env file only in local
	if env == "local" {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	viperConfigPath := os.Getenv("VIPER_CONFIG_PATH")
	// Load the environment-specific configuration file
	viper.AddConfigPath(viperConfigPath)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Override DSN from environment variable if set
	if dsn := os.Getenv("DATABASE_DSN"); dsn != "" {
		config.Database.DSN = dsn
	}

	return &config, nil
}
