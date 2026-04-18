package main

import (
	"context"
	"log"
	"main/handlers"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("MONGODB_URI is not set")
	} else if err := testDatabase(uri); err != nil {
		log.Printf("database test failed: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthcheck)

	r.Route("/user", handlers.UserRouter)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func testDatabase(uri string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("database disconnect failed: %v", err)
		}
	}()

	title := "hello world"
	col := client.Database("app").Collection("users")
	if _, err := col.InsertOne(context.TODO(), bson.D{{"title", title}}); err != nil {
		return err
	}

	return nil
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
