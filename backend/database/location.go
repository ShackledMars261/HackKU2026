package database

import (
	"context"
	"main/errors"
	"main/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateLocation(location *models.Location) *models.Location {
	collection := client.Database("app").Collection("locations")

	model := &models.Location{
		ID:            uuid.NewString(),
		Longitude:     location.Longitude,
		Latitude:      location.Latitude,
		Name:          location.Name,
		OverallRating: 0.0,
	}

	collection.InsertOne(context.Background(), model)

	return model
}

func GetLocation(id string) (*models.Location, error) {
	collection := client.Database("app").Collection("locations")

	var model models.Location
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Join(err, errors.ErrUserNotFound)
		}

		return nil, err
	}

	return &model, nil
}

func GetLocations() ([]*models.Location, error) {
	return nil, nil
}
