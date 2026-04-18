package main

import (
	"context"
	"main/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func RequireSession() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")

			sessionId := cookie.Value
			//verify valid uuid
			_, err = uuid.Parse(sessionId)

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			session, err := database.GetSession(sessionId)
			if err != nil || session == nil {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			if session.ExpiresAt.Before(time.Now()) {
				database.DeleteSession(sessionId)
				http.Error(w, "Session expired", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "session", session)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
