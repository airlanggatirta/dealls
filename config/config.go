// config/config.go

package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
	// Add more configuration options as needed
}

// LoadConfig loads the application configuration from a file
func LoadConfig() (*Config, error) {
	var config Config

	// Set default values
	setDefaultConfig()

	// Read config file (config.yaml or config.json or config.toml)
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// setDefaultConfig sets the default configuration values
func setDefaultConfig() {
	viper.SetDefault("Server.Port", "8080")
	viper.SetDefault("Database.Host", "127.0.0.1")
	viper.SetDefault("Database.Port", "5432")
	viper.SetDefault("Database.Username", "")
	viper.SetDefault("Database.Password", "")
	viper.SetDefault("Database.Name", "database_name")
}
