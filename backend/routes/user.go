package routes

import (
	"main/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", getUser)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := database.GetUser(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, user)
}
