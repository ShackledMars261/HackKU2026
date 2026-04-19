package models

type Asset struct {
	ID         string `bson:"_id" json:"id"`
	LocationID string `bson:"locationID" json:"locationID"`
	ObjectID   string `bson:"objectID" json:"objectID"`
	FileName   string `bson:"fileName" json:"fileName"`
}

type AssetResponse struct {
	Path string `json:"path"`
}

type AssetsResponse struct {
	Paths []string `json:"paths"`
}
