package database

import (
	"context"
	"io"
	"main/errors"
	"main/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreatePost(post *models.Post) error {
	collection := client.Database("app").Collection("posts")

	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}

	return nil
}

func GetPost(id string) (*models.Post, error) {
	collection := client.Database("app").Collection("posts")

	var model models.Post
	if err := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&model); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.Join(err, errors.ErrPostNotFound)
		}

		return nil, err
	}

	return &model, nil
}

func GetPostsForLocation(locationId string) ([]*models.Post, error) {
	collection := client.Database("app").Collection("posts")

	opts := options.Find().SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.Background(), bson.D{{"location_id", locationId}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var models []*models.Post

	if err = cursor.All(context.Background(), &models); err != nil {
		return nil, err
	}
	return models, nil

}

func UploadPhotoToGridFS(filename string, file io.Reader) (string, error) {
	bucket := client.Database("app").GridFSBucket()

	objectID, err := bucket.UploadFromStream(context.Background(), filename, file)
	if err != nil {
		return "", err
	}

	return objectID.Hex(), nil
}

func StreamPhotoFromGridFS(fileID string, w io.Writer) error {
	bucket := client.Database("app").GridFSBucket()

	objectID, err := bson.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	_, err = bucket.DownloadToStream(context.Background(), objectID, w)
	return err
}
