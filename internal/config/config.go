package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres-dev")
	viper.SetDefault("DB_PASSWORD", "mysecretpassword")
	viper.SetDefault("DB_NAME", "dev")
	viper.SetDefault("SERVER_PORT", ":8080")

	viper.AutomaticEnv()

	config := &Config{
		DbHost:     viper.GetString("DB_HOST"),
		DbUser:     viper.GetString("DB_USER"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}

	return config, nil
}
