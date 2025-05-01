package data

import (
	"context"
	"fmt"
	"github.com/Muraddddddddd9/ms-database/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo establishes a connection to MongoDB using configuration
// Returns a MongoDB client instance and any connection error
func ConnectMongo() (*mongo.Client, error) {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load database configuration:", err)
		return nil, err
	}

	// Construct MongoDB connection URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin?authSource=%s",
		cfg.DB_USERNAME,    // MongoDB username
		cfg.DB_PASSWORD,    // MongoDB password
		cfg.DB_HOST,        // MongoDB host address
		cfg.DB_PORT,        // MongoDB port
		cfg.DB_AUTH_SOURCE, // Authentication database
	)

	// Create new MongoDB client with connection options
	client, err := mongo.Connect(
		context.TODO(),                 // Background context
		options.Client().ApplyURI(uri), // Connection URI
	)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return nil, err
	}

	// Verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	return client, nil
}
