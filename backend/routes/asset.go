package routes

import (
	"main/errors"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AssetRouter(r chi.Router) {
	r.Route("/asset", func(r chi.Router) {
		r.With(RequireSession).Post("/{id}", createAsset) // id is location id
		r.Get("/{id}", downloadAsset)

		r.Get("/all/{id}", getAllAssets)
	})
}

func createAsset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	if r.Header.Get("Content-Type") != "application/octet-stream" {
		writeError(w, errors.ErrInvalidBody)
		return
	}

	fileName := r.Header.Get("X-FileName")
	if fileName == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	asset, err := service.UploadAsset(id, fileName, r.Body)
	if err != nil {
		writeError(w, err)
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

func getAllAssets(w http.ResponseWriter, r *http.Request) {
	locationID := chi.URLParam(r, "id")
	if locationID == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	assets, err := service.GetAllAssets(locationID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, assets)
}

func downloadAsset(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	err := service.DownloadAsset(id, w)
	if err != nil {
		writeError(w, err)
		return
	}
}
