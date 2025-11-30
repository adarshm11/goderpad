package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"goderpad/config"
)

var MongoClient *mongo.Client

// InitMongoClient initializes the global MongoClient variable
func InitMongoClient() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI, err := config.GetEnv("MONGO_URI")
	if err != nil {
		log.Printf("Failed to get MONGO_URI from environment variables: %v", err)
		return err
	}

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return err
	}

	MongoClient = client
	log.Printf("MongoClient initialized successfully")
	return nil
}

// ShutdownMongoClient disconnects the global MongoClient
func ShutdownMongoClient() error {
	if MongoClient == nil {
		return fmt.Errorf("MongoClient is not initialized")
	}

	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		log.Printf("Failed to disconnect MongoDB client: %v", err)
		return err
	}

	MongoClient = nil
	log.Printf("MongoClient shut down successfully")
	return nil
}

// TestConnection tests the connection to the MongoDB database -> not for use in production
func TestConnection() error {
	// Ensure the global MongoClient is initialized
	if MongoClient == nil {
		return fmt.Errorf("MongoClient is not initialized")
	}

	// Send a ping to confirm a successful connection
	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return err
	}

	log.Printf("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}
