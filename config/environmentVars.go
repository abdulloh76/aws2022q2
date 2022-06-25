package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	EnvVariables ConfigVariables
)

type ConfigVariables struct {
	PORT                string
	POSTGRES_URI        string
	DOCKER_POSTGRES_URI string
}

func GetEnvConfigs() *ConfigVariables {
	return &EnvVariables
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	err = viper.Unmarshal(&EnvVariables)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}
