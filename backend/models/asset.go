package models

type Asset struct {
	ID       string `bson:"_id" json:"id"`
	ObjectID string `bson:"objectID" json:"objectID"`
	FileName string `bson:"fileName" json:"fileName"`
}

type AssetResponse struct {
	Path string `json:"path"`
}
