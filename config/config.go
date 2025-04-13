package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT              string
	GIN_MODE          string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_URI      string
}

var Envs *EnvConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file error ")
	}

	Envs = &EnvConfig{
		PORT:              getEnv("PORT", "8080"),
		GIN_MODE:          getEnv("GIN_MODE", "debug"),
		POSTGRES_USER:     getEnv("POSTGRES_USER", "postgres"),
		POSTGRES_PASSWORD: getEnv("POSTGRES_PASSWORD", "postgres123"),
		POSTGRES_DB:       getEnv("POSTGRES_DB", "goswechallengedb"),
		POSTGRES_URI:      getEnv("POSTGRES_URI", "postgres:5432"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
