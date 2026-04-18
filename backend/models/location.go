package models

type Location struct {
	ID            string  `bson:"_id" json:"id"`
	Longitude     float64 `bson:"longitude" json:"longitude"`
	Latitude      float64 `bson:"latitude" json:"latitude"`
	Name          string  `bson:"name" json:"name"`
	OverallRating float64 `bson:"overall_rating" json:"overall_rating"`
}
