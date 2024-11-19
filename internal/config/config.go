package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

func InitConfig() error {
	err := InitEnv()
	if err != nil {
		return err
	}

	viper.AddConfigPath("internal/config")
	viper.SetConfigName(".env.local")
	viper.AutomaticEnv()

	return nil
}

func InitEnv() error {
	return godotenv.Load(".env.local")
}
