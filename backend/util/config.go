package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_DSN  string `mapstructure:"POSTGRES_DSN"`
	PORT          string `mapstructure:"PORT"`
	MIGRATION_URL string `mapstructure:"MIGRATION_URL"`
}

func LoadEnv(path string) *Config {
	godotenv.Load(path + "/.env")

	return &Config{
		POSTGRES_DSN:  os.Getenv("POSTGRES_DSN"),
		PORT:          os.Getenv("PORT"),
		MIGRATION_URL: os.Getenv("MIGRATION_URL"),
	}
}
