package service

import (
	"main/database"
	"main/models"

	"github.com/google/uuid"
)

func CreateLocation(session *models.Session, request *models.CreateLocationRequest) (*models.Location, error) {
	location := &models.Location{
		ID: uuid.NewString(),
		Location: models.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{request.Longitude, request.Latitude},
		},
		Name:  request.Name,
		Owner: session.UserID,
	}

	database.InsertLocation(location)

	return location, nil
}

func GetLocation(id string) (*models.Location, error) {
	return database.GetLocation(id)
}

func GetAllLocations() ([]*models.Location, error) {
	return database.GetAllLocations()
}

func GetLocationsNear(request *models.GetLocationsNearRequest) ([]*models.Location, error) {
	return database.GetLocationsNear(request.Longitude, request.Latitude, request.MaxDistanceNear)
}
