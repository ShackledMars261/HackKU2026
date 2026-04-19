package models

import "time"

type StatusReport struct {
	ID         string    `bson:"_id" json:"id"`
	LocationID string    `bson:"location_id" json:"location_id"`
	UserID     string    `bson:"user_id" json:"user_id"`
	Busyness   float64   `bson:"busyness" json:"busyness"`
	Loudness   float64   `bson:"loudness" json:"loudness"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}

type CreateStatusReportRequest struct {
	LocationID string  `json:"location_id"`
	Busyness   float64 `json:"busyness"`
	Loudness   float64 `json:"loudness"`
}

type StatusRequest struct {
	LocationID string
	Recency    *time.Duration
}

type StatusResponse struct {
	LocationID  string  `json: "location_id"`
	AvgBusyness float64 `json: "average_busyness"`
	AvgLoudness float64 `json: "average_loudness"`
}
