package database

import (
	"context"
	"io"
	"main/errors"
	"main/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UploadAsset(locationID string, fileName string, r io.Reader) (*models.Asset, error) {
	bucket := client.Database("app").GridFSBucket(options.GridFSBucket().SetName("asset"))
	collection := client.Database("app").Collection("asset")

	objectID, err := bucket.UploadFromStream(context.Background(), fileName, r)
	if err != nil {
		return nil, err
	}

	asset := &models.Asset{
		ID:         uuid.NewString(),
		LocationID: locationID,
		ObjectID:   objectID.Hex(),
		FileName:   fileName,
	}
	if _, err = collection.InsertOne(context.Background(), asset); err != nil {
		return nil, err
	}

	return asset, nil
}

func GetAsset(id string) (*models.Asset, error) {
	collection := client.Database("app").Collection("asset")

	var asset models.Asset
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&asset); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.ErrAssetNotFound
		}

		return nil, err
	}

	return &asset, nil
}

func GetAllAssets(locationID string) ([]models.Asset, error) {
	collection := client.Database("app").Collection("asset")

	cursor, err := collection.Find(context.Background(), bson.D{{"locationID", locationID}})
	if err != nil {
		return nil, err
	}

	var assets []models.Asset
	if err := cursor.All(context.Background(), &assets); err != nil {
		return nil, err
	}

	return assets, nil
}

func DownloadAsset(id string, w io.Writer) error {
	asset, err := GetAsset(id)
	if err != nil {
		return err
	}

	bucket := client.Database("app").GridFSBucket(options.GridFSBucket().SetName("asset"))

	objectID, err := bson.ObjectIDFromHex(asset.ObjectID)
	if err != nil {
		return errors.ErrAssetNotFound
	}

	_, err = bucket.DownloadToStream(context.Background(), objectID, w)

	return err
}
