package database

import (
	"context"
	"main/errors"
	"main/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InsertUser(user *models.User) error {
	collection := client.Database("app").Collection("user")

	if _, err := collection.InsertOne(context.Background(), user); err != nil {
		return err
	}

	return nil
}

func GetUser(id string) (*models.User, error) {
	collection := client.Database("app").Collection("user")

	var model models.User
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.ErrUserNotFound
		}

		return nil, err
	}

	return &model, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	collection := client.Database("app").Collection("user")

	var model models.User
	if err := collection.FindOne(context.Background(), bson.D{{"username", username}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.ErrUserNotFound
		}

		return nil, err
	}

	return &model, nil
}
