package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is the global MongoDB database instance
var DB *mongo.Database

// LoadConfig loads environment variables and initializes the MongoDB connection.
func LoadConfig() {
	// Attempt to load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading configuration from ENV")
	}

	// Get MongoDB URI from env; fallback to a default value.
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Create client options and connect
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to ensure a successful connection
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Get the database name from env or use "tododb" as default
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "tododb"
	}

	DB = client.Database(dbName)
	log.Println("Connected to MongoDB!")
}
