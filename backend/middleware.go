package main

import (
	"context"
	"main/database"
	"main/errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "invalid session id", http.StatusUnauthorized)
			return
		}

		id := cookie.Value
		if _, err := uuid.Parse(id); err != nil {
			http.Error(w, "invalid session id", http.StatusUnauthorized)
			return
		}

		session, err := database.GetSession(id)
		if err != nil {
			if errors.Is(err, errors.ErrSessionNotFound) {
				http.Error(w, "invalid session id", http.StatusUnauthorized)
			} else {
				http.Error(w, "an unexpected error occured", http.StatusInternalServerError)
			}
			return
		}

		if time.Now().After(session.ExpiresAt) {
			http.Error(w, "session expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "session", session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
