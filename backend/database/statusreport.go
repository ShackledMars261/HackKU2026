package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SaveStatusReport(location *models.StatusReport) error {
	collection := client.Database("app").Collection("statusreports")

	_, err := collection.InsertOne(context.Background(), location)
	if err != nil {
		return err
	}

	return nil
}

func GetStatusReportsForLocationWithRecency(locationId string, recency time.Duration) ([]*models.StatusReport, error) {
	collection := client.Database("app").Collection("statusreports")
	now := time.Now()
	cutoff := now.Add(-recency)

	filter := bson.D{
		{"location_id", locationId},
		{"created_at", bson.D{
			{"$gte", cutoff},
			{"$lte", now},
		}},
	}

	opts := options.Find().SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var models []*models.StatusReport

	if err = cursor.All(context.Background(), &models); err != nil {
		return nil, err
	}
	return models, nil

}

func GetStatusReportsForLocation(locationId string) ([]*models.StatusReport, error) {
	collection := client.Database("app").Collection("statusreports")

	filter := bson.D{
		{"location_id", locationId},
	}

	opts := options.Find().SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var models []*models.StatusReport

	if err = cursor.All(context.Background(), &models); err != nil {
		return nil, err
	}
	return models, nil

}
