package handlers

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
	r.Post("/", createLocation)
	r.Get("/{id}", getLocation)
	//r.Get("/", getAllLocations)
	//r.Get("/search", searchLocations)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var request createLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	new_location := &models.Location{
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		Name:      request.Name,
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
