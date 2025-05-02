package mongodb

import (
	"context"
	"fmt"
	"github.com/Muraddddddddd9/ms-database/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load database configuration:", err)
		return nil, err
	}

	if cfg.MONGO_HOST == "" || cfg.MONGO_PORT == 0 {
		return nil, fmt.Errorf("MongoDB host or port are required")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/admin?authSource=%s",
		cfg.MONGO_USERNAME,
		cfg.MONGO_PASSWORD,
		cfg.MONGO_HOST,
		cfg.MONGO_PORT,
		cfg.MONGO_AUTH_SOURCE,
	)

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return nil, err
	}

	db := client.Database(cfg.MONGO_NAME)

	log.Println("Successfully connected to MongoDB!")
	return db, nil
}
