package errors

import "net/http"

var ErrLocationNotFound = New("location_not_found", "location not found", http.StatusNotFound)
