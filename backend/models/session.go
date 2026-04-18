package models

import (
	"time"
)

type Session struct {
	ID        string    `bson:"_id" json:"id"`
	UserID    string    `bson:"userId" json:"userId"`
	ExpiresAt time.Time `bson:"expiresAt" json:"expiresAt"`
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
