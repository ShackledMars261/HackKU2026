package routes

import (
	"encoding/json"
	"main/database"
	"main/errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteUsers(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", getUser)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := database.GetUser(id)
	if err != nil {
		if errors.Is(err, errors.ErrUserNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, "an unknown error occured", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
