package models

import (
	"time"

	"github.com/google/uuid"
)

var timeout, _ = time.ParseDuration("2m")

type Session struct {
	ID        string    `bson:"_id" json:"id"`
	UserID    string    `bson:"userId" json:"userId"`
	ExpiresAt time.Time `bson:"expiresAt" json:"expiresAt"`
}

func NewSession(userId string) *Session {
	return &Session{
		ID:        uuid.NewString(),
		UserID:    userId,
		ExpiresAt: time.Now().Add(timeout),
	}
}

type SessionStatusResponse struct {
	Exists  bool   `json:"exists"`
	Expired bool   `json:"expired"`
	UserID  string `json:"userId"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
