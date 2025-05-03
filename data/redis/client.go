package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/Muraddddddddd9/ms-database/config"
	"github.com/redis/go-redis/v9"
)

func Connect() (*redis.Client, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
		return nil, err
	}

	if cfg.REDIS_HOST == "" || cfg.REDIS_PORT == 0 {
		return nil, fmt.Errorf("Redis host or port are required")
	}

	addr := fmt.Sprintf("%s:%d", cfg.REDIS_HOST, cfg.REDIS_PORT)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to ping Redis:", err)
		return nil, err
	}

	return client, nil
}