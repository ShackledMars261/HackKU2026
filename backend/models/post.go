package models

type Post struct {
	ID           string   `bson:"_id" json:"id"`
	UserId       string   `bson:"user_id" json:"user_id"`
	LocationId   string   `bson:"location_id" json:"location_id"`
	Rating       float64  `bson:"rating" json:"rating"`
	Description  string   `bson:"description" json:"description"`
	PhotoFileIds []string `bson:"photo_file_ids" json:"photo_file_ids"`
}
