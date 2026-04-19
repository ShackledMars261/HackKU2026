package models

type Location struct {
	ID            string  `bson:"_id" json:"id"`
	Location      GeoJSON `bson:"location" json:"location"`
	Name          string  `bson:"name" json:"name"`
	Owner         string  `bson:"owner" json:"owner"`
	OverallRating float64 `bson:"overall_rating" json:"overallRating"`
}

type NearbyLocation struct {
	Location `bson:",inline"`
	Distance float64 `bson:"distance" json:"distance"`
}

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`               // Must be "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

type CreateLocationRequest struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type GetLocationsNearRequest struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Radius    float64 `json:"radius"`
}
