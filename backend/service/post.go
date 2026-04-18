package service

import (
	"io"
	"main/database"
	"main/models"
	"time"

	"github.com/google/uuid"
)

func CreatePost(session *models.Session, request *models.CreatePostRequest) (*models.Post, error) {
	post := &models.Post{
		ID:           uuid.NewString(),
		UserId:       session.UserID,
		LocationId:   request.LocationId,
		Rating:       request.Rating,
		Description:  request.Description,
		CreatedAt:    time.Now(),
		PhotoFileIds: request.PhotoFileIds,
	}
	err := database.CreatePost(post)
	if err != nil {
		return nil, err
	}

	avg, err := CalculateRatingAverage(post.LocationId)
	if err != nil {
		return nil, err
	}
	database.UpdateLocationRating(post.LocationId, avg)
	return post, nil
}

func CalculateRatingAverage(locationId string) (float64, error) {
	posts, err := database.GetPostsForLocation(locationId)
	if err != nil {
		return 0.0, err
	}
	tot := 0.0
	for _, post := range posts {
		tot += post.Rating
	}
	return tot / float64(len(posts)), err

}

func GetPost(id string) (*models.Post, error) {
	return database.GetPost(id)
}

func GetPostsForLocation(locationId string) ([]*models.Post, error) {
	return database.GetPostsForLocation(locationId)
}

func UploadPhotoToGridFS(filename string, file io.Reader) (string, error) {
	return database.UploadPhotoToGridFS(filename, file)
}

func StreamPhotoFromGridFS(id string, w io.Writer) error {
	return database.StreamPhotoFromGridFS(id, w)
}
