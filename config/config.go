// config/config.go

package config

import (
	"github.com/spf13/viper"
)

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
}

func LoadConfig() (*Config, error) {
	var config Config

	setDefaultConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaultConfig() {
	viper.SetDefault("Server.Port", "8080")
	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", "3306")
	viper.SetDefault("Database.Username", "root")
	viper.SetDefault("Database.Password", "")
	viper.SetDefault("Database.Name", "dealls_test")
}
