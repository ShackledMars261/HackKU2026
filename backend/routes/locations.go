package routes

import (
	"encoding/json"
	"log"
	"main/database"
	"main/errors"
	"main/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type createLocationRequest struct {
	Longitude float64
	Latitude  float64
	Name      string
}

type getLocationNearRequest struct {
	Longitude       float64
	Latitude        float64
	MaxDistanceNear float64
}

func LocationRouter(r chi.Router) {
	r.Post("/location", createLocation)
	r.Get("/location/{id}", getLocation)
	r.Get("/locations/all", getAllLocations)
	r.Get("/locations/near", getLocationsNear)
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

	created_location, err := database.CreateLocation(new_location)
	if err != nil {
		log.Printf("Error creating location: %v", err)
		http.Error(w, "Error creating location", http.StatusInternalServerError)
	}

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

func getLocationsNear(w http.ResponseWriter, r *http.Request) {
	qLon := r.URL.Query().Get("longitude")
	qLat := r.URL.Query().Get("latitude")
	qMax := r.URL.Query().Get("maxDist")

	if qLon == "" || qLat == "" {
		http.Error(w, "Missing query params", http.StatusBadRequest)
		return
	}
	if qMax == "" {
		qMax = "8000" // 8km ~ 5miles
	}
	maxDist, err := strconv.ParseFloat(qMax, 64)
	longitude, err := strconv.ParseFloat(qLon, 64)
	latitude, err := strconv.ParseFloat(qLat, 64)
	if err != nil {
		http.Error(w, "Invalid query params", http.StatusBadRequest)
		return
	}

	locations, err := database.GetLocationsNear(longitude, latitude, maxDist)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			http.Error(w, "No locations found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error getting locations", http.StatusInternalServerError)
		log.Printf("Error getting locations near: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(locations)
}
