package routes

import (
	"encoding/json"
	"main/errors"
	"main/models"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteAccount(r chi.Router) {
	r.Post("/signup", signup)
	r.Post("/signin", signin)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var request models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	user, err := service.Signup(&request)
	if err != nil {
		if errors.Is(err, errors.ErrUsernameAlreadyExists) {
			http.Error(w, "username already exists", http.StatusBadRequest)
		} else {
			http.Error(w, "an unknown error occured", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func signin(w http.ResponseWriter, r *http.Request) {
	var request models.SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	session, err := service.Signin(&request)
	if err != nil {
		if errors.Is(err, errors.ErrInvalidUsernameOrPassword) {
			http.Error(w, "invalid username or password", http.StatusBadRequest)
		} else {
			http.Error(w, "an unknown error occured", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)
}
