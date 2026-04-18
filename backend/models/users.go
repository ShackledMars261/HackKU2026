package models

type User struct {
	ID    string `bson:"_id" json:"id"`
	Email string `bson:"email" json:"email"`
}

type NewUser struct {
	Email string `bson:"email" json:"email"`
}
