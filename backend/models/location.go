package models

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`               // Must be "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

type CreateLocationRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
}

type GetLocationsNearRequest struct {
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	MaxDistanceNear float64 `json:"maxDistanceNear"`
}

type Location struct {
	ID            string  `bson:"_id" json:"id"`
	Location      GeoJSON `bson:"location" json:"location"`
	Name          string  `bson:"name" json:"name"`
	Owner         string  `bson:"owner" json:"owner"`
	OverallRating float64 `bson:"overall_rating" json:"overall_rating"`
}
