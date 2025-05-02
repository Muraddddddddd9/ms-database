package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/Muraddddddddd9/ms-database/config"

	"github.com/gofiber/storage/redis/v3"
)

func Connect() (*redis.Storage, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
		return nil, err
	}

	if cfg.REDIS_HOST == "" || cfg.REDIS_PORT == 0 {
		return nil, fmt.Errorf("Redis host or port are required")
	}

	storage := redis.New(redis.Config{
		Host:     cfg.REDIS_HOST,
		Port:     cfg.REDIS_PORT,
		Password: cfg.REDIS_PASSWORD,
		Database: cfg.REDIS_DB,
	})

	if _, err := storage.Conn().Ping(context.Background()).Result(); err != nil {
		log.Fatal("Failed to ping Redis:", err)
		return nil, err
	}

	return storage, nil
}
