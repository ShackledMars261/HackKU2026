package routes

import (
	"main/errors"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AssetRouter(r chi.Router) {
	r.With(RequireSession).Route("/asset", func(r chi.Router) {
		r.Post("/", createAsset)
		r.Post("/{id}", getAsset)
		r.Get("/{id}", downloadAsset)
	})
}

func createAsset(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/octet-stream" {
		writeError(w, errors.ErrInvalidBody)
		return
	}

	filename := r.Header.Get("X-Filename")
	if filename == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	asset, err := service.UploadAsset(filename, r.Body)
	if err != nil {
		writeError(w, errors.ErrInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, asset)
}

func getAsset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	asset, err := service.GetAsset(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, asset)
}

func downloadAsset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/jpeg")
	err := service.DownloadAsset(id, w)
	if err != nil {
		writeError(w, err)
		return
	}
}
