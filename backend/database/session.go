package database

import (
	"context"
	"main/models"

	"github.com/google/uuid"
)

func CreateSession(session *models.Session) error {
	collection := client.Database("app").Collection("session")

	if _, err := collection.InsertOne(context.Background(), session); err != nil {
		return err
	}

	return nil
}

func GetSession(sessionId uuid.UUID) (*models.Session, error) {
	return nil, nil
}

func DeleteSession(sessionId uuid.UUID) error {
	return nil
}
