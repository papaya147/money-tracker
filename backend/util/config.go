package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DSN string `mapstructure:"POSTGRES_DSN"`
	PORT         string `mapstructure:"PORT"`
}

func LoadEnv(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("error loading env vars:", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Panic("error unmarshalling config:", err)
	}

	return &config
}
