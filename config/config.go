package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration parameters
// Contains settings for database, Redis, and admin credentials
type Config struct {
	// Database configuration
	DB_HOST        string `env:"DB_HOST"`        // MongoDB host address
	DB_PORT        string `env:"DB_PORT"`        // MongoDB port number
	DB_USERNAME    string `env:"DB_USERNAME"`    // MongoDB username
	DB_PASSWORD    string `env:"DB_PASSWORD"`    // MongoDB password
	DB_AUTH_SOURCE string `env:"DB_AUTH_SOURCE"` // MongoDB authentication database

	// Redis configuration
	REDIS_HOST     string `env:"REDIS_HOST"`     // Redis server host
	REDIS_PORT     string `env:"REDIS_PORT"`     // Redis server port
	REDIS_PASSWORD string `env:"REDIS_PASSWORD"` // Redis authentication password
}

// LoadConfig loads configuration from .env file and environment variables
// Returns a Config struct populated with values or an error if loading fails
//
// Example:
// cfg, err := LoadConfig()
//
//	if err != nil {
//	    log.Fatal("Failed to load config:", err)
//	}
func LoadConfig() (*Config, error) {
	// Try loading .env file (ignored if not present)
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Initialize config with environment variables
	cfg := &Config{
		DB_HOST:        GetEnv("DB_HOST"),
		DB_PORT:        GetEnv("DB_PORT"),
		DB_USERNAME:    GetEnv("DB_USERNAME"),
		DB_PASSWORD:    GetEnv("DB_PASSWORD"),
		DB_AUTH_SOURCE: GetEnv("DB_AUTH_SOURCE"),

		REDIS_HOST:     GetEnv("REDIS_HOST"),
		REDIS_PORT:     GetEnv("REDIS_PORT"),
		REDIS_PASSWORD: GetEnv("REDIS_PASSWORD"),
	}

	// Validate required fields
	if cfg.DB_HOST == "" || cfg.DB_PORT == "" {
		return nil, fmt.Errorf("database host and port are required")
	}

	return cfg, nil
}

// GetEnv is a helper function to get environment variables with empty string fallback
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return "" // Return empty string if not set
	}
	return value
}
