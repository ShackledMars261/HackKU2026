package models

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`               // Must be "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

type Location struct {
	ID            string  `bson:"_id" json:"id"`
	Location      GeoJSON `bson:"location" json:"location"`
	Name          string  `bson:"name" json:"name"`
	OverallRating float64 `bson:"overall_rating" json:"overall_rating"`
}
