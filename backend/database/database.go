package database

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Initialize() {
	uri := os.Getenv("MONGODB_URI")

	client, _ = mongo.Connect(options.Client().ApplyURI(uri))
}
