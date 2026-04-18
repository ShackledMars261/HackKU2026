package models

import "github.com/google/uuid"

type User struct {
	ID    string `bson:"_id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

func NewUser(name, email string) User {
	return User{
		ID:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
}
