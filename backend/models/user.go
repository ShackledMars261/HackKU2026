package models

type User struct {
	ID       string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password []byte `bson:"password" json:"-"`
}
