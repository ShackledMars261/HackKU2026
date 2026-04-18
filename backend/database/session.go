package database

import (
	"context"
	"main/errors"
	"main/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateSession(session *models.Session) error {
	collection := client.Database("app").Collection("session")

	if _, err := collection.InsertOne(context.Background(), session); err != nil {
		return err
	}

	return nil
}

func GetSession(sessionId string) (*models.Session, error) {
	collection := client.Database("app").Collection("session")

	var session models.Session
	if err := collection.FindOne(context.Background(), bson.D{{"_id", sessionId}}).Decode(&session); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Join(err, errors.ErrSessionNotFound)
		}
		return nil, err
	}
	return &session, nil
}

func DeleteSession(sessionId string) error {
	return nil
}
