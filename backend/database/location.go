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
	MinDistance float64        `bson:"$minDistance"`
	MaxDistance float64        `bson:"$maxDistance"`
}

func GetNearbyLocations(request *models.GetLocationsNearRequest) ([]models.NearbyLocation, error) {
	collection := client.Database("app").Collection("location")

	aggregation := mongo.Pipeline{
		{{
			"$geoNear",
			bson.D{
				{"near", models.GeoJSON{
					Type:        "Point",
					Coordinates: []float64{request.Longitude, request.Latitude},
				}},
				{"distanceField", "distance"},
				{"distanceMultiplier", 0.00062137273},
				{"maxDistance", request.Radius * 1609.34},
			},
		}},
		{{
			"$sort",
			bson.D{{"distance", 1}}, // asc
		}},
	}

	cursor, err := collection.Aggregate(context.Background(), aggregation)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	locations := []models.NearbyLocation{}
	if err := cursor.All(context.Background(), &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
