package service

import (
	"io"
	"main/database"
	"main/models"
)

func UploadAsset(locationID string, fileName string, r io.Reader) (*models.AssetResponse, error) {
	asset, err := database.UploadAsset(locationID, fileName, r)
	if err != nil {
		return nil, err
	}

	return &models.AssetResponse{
		Path: "/asset/" + asset.ID,
	}, nil
}

func GetAsset(id string) (*models.Asset, error) {
	return database.GetAsset(id)
}

func DownloadAsset(id string, w io.Writer) error {
	return database.DownloadAsset(id, w)
}
