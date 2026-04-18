package errors

import "net/http"

var ErrPostNotFound = New("post_not_found", "Post not found.", http.StatusNotFound)
var ErrPhotoNotFound = New("photo_not_found", "photo not found", http.StatusNotFound)
var ErrUnauthorized = New("unauthorized", "Unauthorized", http.StatusUnauthorized)
var ErrBadFormData = New("bad_form_data", "Unable to parse form data", http.StatusBadRequest)
var ErrBadRatingFormat = New("bad_rating_fomrat", "Invalid rating format", http.StatusBadRequest)
var ErrMissingParam = New("missing_url_param", "Missing url parameter", http.StatusBadRequest)
var ErrInternalServerError = New("internal_server_error", "Internal server error", http.StatusInternalServerError)
