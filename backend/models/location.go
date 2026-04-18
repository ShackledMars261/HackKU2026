package models

type Location struct {
	ID            string  `bson:"_id" json:"id"`
	Location      GeoJSON `bson:"location" json:"location"`
	Name          string  `bson:"name" json:"name"`
	Owner         string  `bson:"owner" json:"owner"`
	OverallRating float32 `bson:"overall_rating" json:"overallRating"`
}

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`               // Must be "Point"
	Coordinates []float32 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

type CreateLocationRequest struct {
	Name      string  `json:"name"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type GetLocationsNearRequest struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Radius    float32 `json:"radius"`
}
