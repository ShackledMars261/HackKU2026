package errors

import "net/http"

var (
	ErrInvalidSessionID = New("invalid_session", "invalid session id", http.StatusUnauthorized)
	ErrSessionNotFound  = New("session_not_found", "session not found", http.StatusNotFound)
	ErrSessionExpired   = New("session_expired", "session expired", http.StatusUnauthorized)
)
