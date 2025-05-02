package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type GlobalConfig struct {
	MONGO_NAME        string `env:"MONGO_NAME"`
	MONGO_HOST        string `env:"MONGO_HOST"`
	MONGO_PORT        int    `env:"MONGO_PORT"`
	MONGO_USERNAME    string `env:"MONGO_USERNAME"`
	MONGO_PASSWORD    string `env:"MONGO_PASSWORD"`
	MONGO_AUTH_SOURCE string `env:"MONGO_AUTH_SOURCE"`

	REDIS_DB       int    `env:"REDIS_HOST"`
	REDIS_HOST     string `env:"REDIS_HOST"`
	REDIS_PORT     int    `env:"REDIS_PORT"`
	REDIS_PASSWORD string `env:"REDIS_PASSWORD"`
}

func LoadConfig() (*GlobalConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &GlobalConfig{
		MONGO_NAME:        GetEnv("MONGO_NAME"),
		MONGO_HOST:        GetEnv("MONGO_HOST"),
		MONGO_PORT:        GetEnvInt("MONGO_PORT"),
		MONGO_USERNAME:    GetEnv("MONGO_USERNAME"),
		MONGO_PASSWORD:    GetEnv("MONGO_PASSWORD"),
		MONGO_AUTH_SOURCE: GetEnv("MONGO_AUTH_SOURCE"),

		REDIS_DB:       GetEnvInt("REDIS_DB"),
		REDIS_HOST:     GetEnv("REDIS_HOST"),
		REDIS_PORT:     GetEnvInt("REDIS_PORT"),
		REDIS_PASSWORD: GetEnv("REDIS_PASSWORD"),
	}

	return cfg, nil
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}

func GetEnvInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		return 0
	}
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int(intValue)
}
