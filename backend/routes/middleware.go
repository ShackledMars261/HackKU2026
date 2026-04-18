package routes

import (
	"context"
	"main/database"
	"main/errors"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeError(w, errors.ErrInvalidSessionID)
			return
		}

		id := authorization[7:]
		if _, err := uuid.Parse(id); err != nil {
			writeError(w, errors.ErrInvalidSessionID)
			return
		}

		session, err := database.GetSession(id)
		if err != nil {
			writeError(w, err)
			return
		}

		if time.Now().After(session.ExpiresAt) {
			writeError(w, errors.ErrSessionExpired)
			return
		}

		ctx := context.WithValue(r.Context(), "session", session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
