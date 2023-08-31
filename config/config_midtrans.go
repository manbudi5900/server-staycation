package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Database DatabaseConfigurations
	Midtrans PaymentConfigurations
}

// DatabaseConfigurations
type DatabaseConfigurations struct {
	Host       string `mapstructure:"Host"`
	Port       string `mapstructure:"Port"`
	DBName     string `mapstructure:"DBName"`
	DBUser     string `mapstructure:"DBUser"`
	DBPassword string `mapstructure:"DBPassword"`
}

// PaymentConfigurations
type PaymentConfigurations struct {
	ClientKey string `mapstructure:"ClientKey"`
	ServerKey string `mapstructure:"ServerKey"`
	APIEnv    string `mapstructure:"APIEnv"`
}

func LoadConfig(path string) (config MainConfig, err error) {
	

	viper.SetConfigFile("config.yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		
		return config, err
	}
	

	err = viper.Unmarshal(&config.Midtrans)
	if err != nil {
		fmt.Printf("Failed to unmarshal")
		return config, err
	}
	return config, nil
}