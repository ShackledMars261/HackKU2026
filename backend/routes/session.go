package routes

import (
	"main/models"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SessionRouter(r chi.Router) {
	r.Post("/signup", signup)
	r.Post("/signin", signin)
	r.Get("/session/{id}", session)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var request models.SignupRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	session, err := service.Signup(&request)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, session)
}

func signin(w http.ResponseWriter, r *http.Request) {
	var request models.SigninRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	session, err := service.Signin(&request)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, session)
}

func session(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	status, err := service.GetSessionStatus(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, status)
}
