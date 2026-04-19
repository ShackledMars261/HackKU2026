package main

import (
	"log"
	"main/database"
	"main/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.Initialize()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthcheck)

	routes.LocationRouter(r)
	routes.UserRouter(r)
	routes.SessionRouter(r)
	routes.StatusReportRouter(r)
	routes.AssetRouter(r)
	routes.PostRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
