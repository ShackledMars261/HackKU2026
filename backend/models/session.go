package models

import (
	"time"

	"github.com/google/uuid"
)

var timeout, _ = time.ParseDuration("2m")

type Session struct {
	ID        string    `bson:"_id" json:"id"`
	UserId    string    `bson:"user_id" json:"user_id"`
	ExpiresAt time.Time `bson:"expires_at" json:"expires_at"`
}

func NewSession(userId string) *Session {
	return &Session{
		ID:        uuid.NewString(),
		UserId:    userId,
		ExpiresAt: time.Now().Add(timeout),
	}
}
