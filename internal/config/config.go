package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var envs = []string{"RABBITMQ_USER", "RABBITMQ_PASSWORD"}

type Config struct {
	BrokerUser string `mapstructure:"RABBITMQ_USER"`
	BrokerPass string `mapstructure:"RABBITMQ_PASSWORD"`
}

func InitConfig() (Config, error) {
	var cfg Config

	err := InitEnv()
	if err != nil {
		return cfg, err
	}

	viper.AddConfigPath("internal/config")
	viper.SetConfigName(".env.local")

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func InitEnv() error {
	return godotenv.Load(".env.local")
}
