package errors

import "net/http"

var (
	ErrInvalidBody = New("invalid_body", "request body is invalid or malformed", http.StatusBadRequest)
	ErrUnknown     = New("internal_error", "an unexpected error occurred", http.StatusInternalServerError)
)
