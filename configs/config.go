package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var AppConfig Config

func loadConfigFile(fileConfigName string) error {
	viper.SetConfigName(fmt.Sprintf("config.%s", fileConfigName))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		return err
	}

	return nil
}

func LoadConfig(fileConfigName string) error {
	if err := loadConfigFile(fileConfigName); err != nil {
		return err
	}

	if fileConfigName != "default" {
		viper.SetConfigName(fmt.Sprintf("config.%s", fileConfigName))
		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}
