package routes

import (
	"encoding/json"
	"main/database"
	"main/errors"
	"main/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserRouter(r chi.Router) {
	r.Post("/", createUser)
	r.Get("/{id}", getUser)
	// r.Put("/{id}", putUser)
	// r.Delete("/{id}", deleteUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	new_user := &models.NewUser{
		Email: request.Email,
	}

	created_user := database.CreateUser(new_user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created_user)
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
