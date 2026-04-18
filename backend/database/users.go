package database

import (
	"context"
	"main/errors"
	"main/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateUser(user *models.NewUser) *models.User {
	collection := client.Database("app").Collection("users")

	model := &models.User{
		ID:    uuid.NewString(),
		Email: user.Email,
	}

	collection.InsertOne(context.Background(), model)

	return model
}

func GetUser(id string) (*models.User, error) {
	collection := client.Database("app").Collection("users")

	var model models.User
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Join(err, errors.ErrUserNotFound)
		}

		return nil, err
	}

	return &model, nil
}
