package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthcheck)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", healthcheck)
		r.Get("/{userId}", healthcheck)
		r.Put("/{userId}", healthcheck)
		r.Delete("/{userId}", healthcheck)
	})

	http.ListenAndServe(":8080", r)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
