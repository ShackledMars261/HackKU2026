package errors

import "net/http"

var ErrPostNotFound = New("post_not_found", "Post not found.", http.StatusNotFound)
