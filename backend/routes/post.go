package routes

import (
	"encoding/json"
	"log"
	"main/database"
	"main/errors"
	"main/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RoutePosts(r chi.Router) {
	r.Post("/post", createPost)
	r.Get("/post/{id}", getPost)
	r.Get("/posts/{locationId}", getPostsByLocation)
	r.Get("/photo/{photoId}", getPhoto)
}

func createPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil && err != http.ErrNotMultipart {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	userId := r.FormValue("user_id")
	locationId := r.FormValue("location_id")
	description := r.FormValue("description")

	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		http.Error(w, "Invalid rating format", http.StatusBadRequest)
		return
	}

	photoFileIds := make([]string, 0)

	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		files := r.MultipartForm.File["photos"]

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, "Failed to read uploaded file", http.StatusInternalServerError)
				return
			}

			fileId, err := database.UploadPhotoToGridFS(fileHeader.Filename, file)
			file.Close()

			if err != nil {
				log.Printf("GridFS upload error: %v", err)
				http.Error(w, "Error saving photo to database", http.StatusInternalServerError)
				return
			}

			photoFileIds = append(photoFileIds, fileId)
		}
	}

	new_post := &models.Post{
		UserId:       userId,
		LocationId:   locationId,
		Rating:       rating,
		Description:  description,
		PhotoFileIds: photoFileIds,
	}

	created_post, err := database.CreatePost(new_post)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created_post)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := database.GetPost(id)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "photoId")
	if id == "" {
		http.Error(w, "Invalid photo ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=31536000")

	err := database.StreamPhotoFromGridFS(id, w)
	if err != nil {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}
}

func getPostsByLocation(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "locationId")
	log.Printf("Getting posts by location: %v", locationId)

	if locationId == "" {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	posts, err := database.GetPostsForLocation(locationId)
	if err != nil {
		log.Printf("Error getting posts for location %s: %v", locationId, err)
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
