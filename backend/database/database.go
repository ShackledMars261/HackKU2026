package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Initialize() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("MONGODB_URI is not set")
	}

	var err error
	client, err = mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	EnsureLocationIndexes()
}

func EnsureLocationIndexes() {
	collection := client.Database("app").Collection("location")

	indexModel := mongo.IndexModel{
		Keys: bson.D{{"location", "2dsphere"}},
	}

	name, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create geospatial index: %v", err)
	}

	log.Printf("Ensured index exists: %s", name)
}
