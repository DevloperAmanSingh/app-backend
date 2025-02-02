package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbName string
var userCollection *mongo.Collection
var tripCollection *mongo.Collection

func ConnectDatabase(url, name string) {
	dbName = name
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	userCollection = client.Database(dbName).Collection("users")
	tripCollection = client.Database(dbName).Collection("trips")

	// Create geospatial index on the "location" field
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	}
	_, err = tripCollection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Error creating geospatial index: %v", err)
	}

	log.Println("Connected to MongoDB!")
}

func GetUserCollection() *mongo.Collection {
	if userCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return userCollection
}

func GetTripCollection() *mongo.Collection {
	if tripCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return tripCollection
}

func DisconnectDatabase() {
	if client != nil {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from database: %v", err)
		}
	}
}
