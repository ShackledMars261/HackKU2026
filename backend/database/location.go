package database

import (
	"context"
	"main/errors"
	"main/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InsertLocation(location *models.Location) error {
	collection := client.Database("app").Collection("location")

	_, err := collection.InsertOne(context.Background(), location)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLocationRating(id string, newRating float64) error {
	collection := client.Database("app").Collection("location")

	filter := bson.D{{"_id", id}}

	update := bson.D{
		{"$set", bson.D{
			{"overall_rating", newRating},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.ErrLocationNotFound
	}

	return nil
}

func GetLocation(id string) (*models.Location, error) {
	collection := client.Database("app").Collection("location")

	var model models.Location
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.ErrLocationNotFound
		}

		return nil, err
	}

	return &model, nil
}

func GetAllLocations() ([]models.Location, error) {
	collection := client.Database("app").Collection("location")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	locations := []models.Location{}

	if err = cursor.All(context.Background(), &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

type nearSphere struct {
	Geometry    models.GeoJSON `bson:"$geometry"`
	MinDistance float32        `bson:"$minDistance"`
	MaxDistance float32        `bson:"$maxDistance"`
}

func GetNearbyLocations(request *models.GetLocationsNearRequest) ([]models.Location, error) {
	collection := client.Database("app").Collection("location")

	filter := bson.D{{
		Key: "location",
		Value: bson.D{{
			Key: "$nearSphere",
			Value: nearSphere{
				Geometry: models.GeoJSON{
					Type:        "Point",
					Coordinates: []float32{request.Longitude, request.Latitude},
				},
				MinDistance: 10,
				MaxDistance: 1000,
			},
		}},
	}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	locations := []models.Location{}
	if err := cursor.All(context.Background(), &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
