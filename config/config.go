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
	SETUP_DB_URL      string
	SETUP_DB_TOKEN    string
	POLYGON_API_KEY   string
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
		POSTGRES_DB:       getEnv("POSTGRES_DB", "postgres"),
		POSTGRES_URI:      getEnv("POSTGRES_URI", "postgres:5432"),
		SETUP_DB_URL:      getEnv("SETUP_DB_URL", ""),
		SETUP_DB_TOKEN:    getEnv("SETUP_DB_TOKEN", ""),
		POLYGON_API_KEY:   getEnv("POLYGON_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
