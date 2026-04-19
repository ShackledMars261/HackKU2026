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

	if err := database.InsertLocation(location); err != nil {
		return nil, err
	}

	return location, nil
}

func GetLocation(id string) (*models.Location, error) {
	return database.GetLocation(id)
}

func GetAllLocations() ([]models.Location, error) {
	return database.GetAllLocations()
}

func GetNearbyLocations(request *models.GetLocationsNearRequest) ([]models.NearbyLocation, error) {
	return database.GetNearbyLocations(request)
}
