package routes

import (
	"main/models"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LocationRouter(r chi.Router) {
	r.With(RequireSession).Route("/location", func(r chi.Router) {
		r.Post("/", createLocation)
		r.Post("/{id}", getLocation)
	})

	r.With(RequireSession).Post("/location/all", getAllLocations)
	r.With(RequireSession).Post("/location/nearby", getNearbyLocations)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var request models.CreateLocationRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	// session will always be valid if RequireSession is used
	session, _ := r.Context().Value("session").(*models.Session)

	location, err := service.CreateLocation(session, &request)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, location)
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

func getNearbyLocations(w http.ResponseWriter, r *http.Request) {
	var request models.GetLocationsNearRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	locations, err := service.GetNearbyLocations(&request)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, locations)
}
