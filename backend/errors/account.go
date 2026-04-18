package errors

import "net/http"

var (
	ErrUsernameAlreadyExists     = New("username_exists", "username already exists", http.StatusConflict)
	ErrInvalidUsernameOrPassword = New("invalid_credentials", "invalid username or password", http.StatusUnauthorized)
)
