package data

import (
	"github.com/Muraddddddddd9/ms-database/config"
	"log"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
)

// ConnectRedis establishes a connection to Redis server using configuration
// Returns a Redis storage instance and any connection error
func ConnectRedis() (*redis.Storage, error) {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
		return nil, err
	}

	// Convert Redis port from string to integer
	port, err := strconv.Atoi(cfg.REDIS_PORT)
	if err != nil {
		log.Fatal("Invalid Redis port format:", err)
		return nil, err
	}

	// Initialize Redis storage with configuration
	storage := redis.New(redis.Config{
		Host:     cfg.REDIS_HOST,     // Redis server hostname
		Port:     port,               // Redis server port
		Password: cfg.REDIS_PASSWORD, // Redis authentication password
		Database: 0,                  // Redis database number (default 0)
	})

	return storage, nil
}
