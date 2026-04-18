package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {
	r.Post("/", postUser)
	r.Get("/{id}", getUser)
	r.Put("/{id}", putUser)
	r.Delete("/{id}", deleteUser)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func putUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
