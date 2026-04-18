package routes

import (
	"encoding/json"
	"main/database"
	"main/errors"
	"main/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type createLocationRequest struct {
	Longitude float64
	Latitude  float64
	Name      string
}

func LocationRouter(r chi.Router) {
	r.Post("/location", createLocation)
	r.Get("/location/{id}", getLocation)
	r.Get("/locations/all", getAllLocations)
	//r.Get("/locations/search", searchLocations)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var request createLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	new_location := &models.Location{
		Location: models.GeoJSON{
			Type:        "Point",
			Coordinates: []float64{request.Longitude, request.Latitude},
		},
		Name: request.Name,
	}

	created_location := database.CreateLocation(new_location)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created_location)
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	location, err := database.GetLocation(id)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			http.Error(w, "Location not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}

func getAllLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := database.GetAllLocations()
	if err != nil {
		http.Error(w, "Error getting locations", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
