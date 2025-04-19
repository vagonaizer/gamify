package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		BaseIp string `mapstructure:"base_ip"` // Base IP address for the server
		Port   string `mapstructure:"port"`    // Port number for the server
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`     // Hostname or IP address of the database server
		Port     string `mapstructure:"port"`     // Port number for the database server
		User     string `mapstructure:"user"`     // Username for the database
		Password string `mapstructure:"password"` // Password for the database user
		Name     string `mapstructure:"name"`     // Name of the database
	} `mapstructure:"database"`
	Logger struct {
		Level string `mapstructure:"level"` // Logging level (e.g., "debug", "info", "warn", "error")
	} `mapstructure:"logger"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath + "config.yaml") // Path to the config file
	viper.AutomaticEnv()                         // Automatically read environment variables

	// values by-default
	viper.SetDefault("server.baseip", "0.0.0.0")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("logger.level", "info")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		// if file not found, return default config
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}
