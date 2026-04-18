package handlers

import (
	"encoding/json"
	"main/database"
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
	// r.Get("/{id}", getUser)
	// r.Put("/{id}", putUser)
	// r.Delete("/{id}", deleteUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.NewUser{
		Email: request.Email,
		Name:  request.Name,
	}

	database.CreateUser(user)
}
