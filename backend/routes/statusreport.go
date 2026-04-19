package routes

import (
	"log"
	"main/errors"
	"main/models"
	"main/service"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func StatusReportRouter(r chi.Router) {
	r.With(RequireSession).Route("/status", func(r chi.Router) {
		r.Post("/", createStatusReport)
		r.Get("/{locationId}", getStatusForLocation)
	})

}

func createStatusReport(w http.ResponseWriter, r *http.Request) {
	var request models.CreateStatusReportRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	session, ok := r.Context().Value("session").(*models.Session)
	if !ok || session == nil {
		writeError(w, errors.ErrInvalidSessionID)
		return
	}

	createdStatusReport, err := service.CreateStatusReport(session, &request)
	if err != nil {
		log.Printf("Error creating status report: %v", err)
		writeError(w, errors.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, createdStatusReport)
}

func getStatusForLocation(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "locationId")
	recencyParam := r.URL.Query().Get("recency")
	log.Printf("Getting status by location: %v", locationId)

	session, ok := r.Context().Value("session").(*models.Session)
	if !ok || session == nil {
		writeError(w, errors.ErrInvalidSessionID)
		return
	}

	if locationId == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	var recency *time.Duration
	if recencyParam != "" {
		parsedDuration, err := time.ParseDuration(recencyParam)
		if err != nil {
			writeError(w, errors.ErrInvalidDurationFormat)
			return
		}

		recency = &parsedDuration
	}

	request := &models.StatusRequest{
		LocationID: locationId,
		Recency:    recency,
	}

	resp, err := service.GetStatusForLocation(session, request)
	if err != nil {
		log.Printf("Error getting status for location %s: %v", locationId, err)
		writeError(w, errors.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
