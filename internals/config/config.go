package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	ConfigName = "default"
	ConfigType = "toml"
	ConfigPath = "./config"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Client   ClientConfig
	External ExternalConfig
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
		return nil, err
	}

	log.Println("Configuration loaded successfully")
	return &config, nil
}
