package database

import (
	"context"
	"main/models"

	"github.com/google/uuid"
)

func CreateSession(userId uuid.UUID) (*models.Session, error) {
	collection := client.Database("app").Collection("session")
	model := models.NewSession(userId)

	collection.InsertOne(context.Background(), model)
	return nil, nil
}

func GetSession(sessionId uuid.UUID) (*models.Session, error) {
	return nil, nil
}

func DeleteSession(sessionId uuid.UUID) error {
	return nil
}
