package database

import (
	"context"
	"main/models"

	"github.com/google/uuid"
)

func CreateUser(user *models.NewUser) {
	collection := client.Database("app").Collection("users")

	model := &models.User{
		ID:    uuid.NewString(),
		Email: user.Email,
		Name:  user.Name,
	}

	collection.InsertOne(context.Background(), model)
}
