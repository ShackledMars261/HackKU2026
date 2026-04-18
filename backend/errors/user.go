package errors

import "net/http"

var (
	ErrUserNotFound = New("user_not_found", "user not found", http.StatusNotFound)
)
