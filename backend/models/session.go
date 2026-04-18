package models

import (
	"time"

	"github.com/google/uuid"
)

var timeout, _ = time.ParseDuration("2m")

type Session struct {
	SessionId string `bson:"_id" json:"id"`
	UserId    string `bson:"userId", json:"userId"`
	ExpiresAt time.Time
}

func NewSession(userId uuid.UUID) *Session {
	return &Session{
		SessionId: uuid.NewString(),
		UserId:    userId.String(),
		ExpiresAt: time.Now().Add(timeout),
	}
}
