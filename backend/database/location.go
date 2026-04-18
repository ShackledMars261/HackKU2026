package database

import (
	"context"
	"main/errors"
	"main/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateLocation(location *models.Location) (*models.Location, error) {
	collection := client.Database("app").Collection("locations")

	model := &models.Location{
		ID:            uuid.NewString(),
		Location:      location.Location,
		Name:          location.Name,
		OverallRating: 0.0,
	}

	_, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return nil, err
	}

	return model, nil
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

func GetAllLocations() ([]*models.Location, error) {
	collection := client.Database("app").Collection("locations")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var locations []*models.Location

	if err = cursor.All(context.Background(), &locations); err != nil {
		return nil, err
	}
	return locations, nil

}

func GetLocationsNear(lon, lat, maxDistMeters float64) ([]*models.Location, error) {
	collection := client.Database("app").Collection("locations")
	filter := bson.D{
		{"location", bson.D{
			{"$near", bson.D{
				{"$geometry", bson.D{
					{"type", "Point"},
					{"coordinates", bson.A{lon, lat}},
				}},
				{"$maxDistance", maxDistMeters},
			}},
		}},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var locations []*models.Location

	if err := cursor.All(context.Background(), &locations); err != nil {
		return nil, err
	}

	if locations == nil {
		locations = []*models.Location{}
	}

	return locations, nil
}
