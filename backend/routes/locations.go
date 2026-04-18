package routes

import (
	"encoding/json"
	"main/errors"
	"main/models"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func LocationRouter(r chi.Router) {
	r.With(RequireSession).Route("/location", func(r chi.Router) {
		r.Post("/", createLocation)
		r.Get("/{id}", getLocation)
	})

	r.With(RequireSession).Get("/locations/all", getAllLocations)
	r.With(RequireSession).Get("/locations/near", getLocationsNear)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var request models.CreateLocationRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	session, ok := r.Context().Value("session").(*models.Session)
	if !ok || session == nil {
		writeError(w, errors.ErrInvalidSessionID)
		return
	}

	createdLocation, err := service.CreateLocation(session, &request)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdLocation)
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	location, err := service.GetLocation(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, location)
}

func getAllLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := service.GetAllLocations()
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, locations)
}

func getLocationsNear(w http.ResponseWriter, r *http.Request) {
	qLon := r.URL.Query().Get("longitude")
	qLat := r.URL.Query().Get("latitude")
	qMax := r.URL.Query().Get("maxDist")

	if qLon == "" || qLat == "" {
		writeError(w, errors.ErrInvalidBody)
		return
	}
	if qMax == "" {
		qMax = "8000"
	}
	maxDist, err := strconv.ParseFloat(qMax, 64)
	if err != nil {
		writeError(w, errors.ErrInvalidBody)
		return
	}
	longitude, err := strconv.ParseFloat(qLon, 64)
	if err != nil {
		writeError(w, errors.ErrInvalidBody)
		return
	}
	latitude, err := strconv.ParseFloat(qLat, 64)
	if err != nil {
		writeError(w, errors.ErrInvalidBody)
		return
	}

	locations, err := service.GetLocationsNear(&models.GetLocationsNearRequest{
		Longitude:       longitude,
		Latitude:        latitude,
		MaxDistanceNear: maxDist,
	})
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, locations)
}
