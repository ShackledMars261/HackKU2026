package routes

import (
	"encoding/json"
	"log"
	"main/errors"
	"main/models"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RoutePosts(r chi.Router) {
	r.With(RequireSession).Route("/post", func(r chi.Router) {
		r.Post("/", createPost)
		r.Get("/{id}", getPost)
	})
	r.With(RequireSession).Post("/posts/nopic", createPostNoPictures)
	r.Get("/posts/{locationId}", getPostsByLocation)
	r.Get("/photo/{photoId}", getPhoto)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("session").(*models.Session)
	if !ok {
		writeError(w, errors.ErrUnauthorized)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil && err != http.ErrNotMultipart {
		writeError(w, errors.ErrBadFormData)
		return
	}

	locationId := r.FormValue("location_id")
	description := r.FormValue("description")

	rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
	if err != nil {
		writeError(w, errors.ErrBadRatingFormat)
		return
	}

	photoFileIds := make([]string, 0)

	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		files := r.MultipartForm.File["photos"]

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				writeError(w, errors.ErrInternalServerError)
				log.Printf("Error reading uploaded file: %v", err)
				return
			}

			fileId, err := service.UploadPhotoToGridFS(fileHeader.Filename, file)
			file.Close()

			if err != nil {
				writeError(w, errors.ErrInternalServerError)
				log.Printf("GridFS upload error: %v", err)
				return
			}

			photoFileIds = append(photoFileIds, fileId)
		}
	}

	postRequest := &models.CreatePostRequest{
		LocationId:   locationId,
		Rating:       rating,
		Description:  description,
		PhotoFileIds: photoFileIds,
	}

	created_post, err := service.CreatePost(session, postRequest)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		writeError(w, errors.ErrInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created_post)
}

func createPostNoPictures(w http.ResponseWriter, r *http.Request) {
	var request models.CreatePostRequest
	if err := readBody(r, &request); err != nil {
		writeError(w, err)
		return
	}

	session, _ := r.Context().Value("session").(*models.Session)
	post, err := service.CreatePost(session, &request)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		writeError(w, errors.ErrInternalServerError)
	}

	writeJSON(w, http.StatusOK, post)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, err := service.GetPost(id)
	if err != nil {
		if errors.Is(err, errors.ErrPostNotFound) {
			writeError(w, errors.ErrPostNotFound)
			return
		} else {
			log.Printf("Error getting post: %v", err)
			writeError(w, errors.ErrInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "photoId")
	if id == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=31536000")

	err := service.StreamPhotoFromGridFS(id, w)
	if err != nil {
		writeError(w, errors.ErrPhotoNotFound)
		return
	}
}

func getPostsByLocation(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "locationId")
	log.Printf("Getting posts by location: %v", locationId)

	if locationId == "" {
		writeError(w, errors.ErrMissingParam)
		return
	}

	posts, err := service.GetPostsForLocation(locationId)
	if err != nil {
		log.Printf("Error getting posts for location %s: %v", locationId, err)
		writeError(w, errors.ErrInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
