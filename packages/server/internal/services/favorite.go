package services

import (
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/database"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/helpers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
)

func GetFavorites(uId uint) ([]models.AssetResponse, error) {
	db := database.GetDBConnection()
	var favorites []models.Favorite
	if err := db.Preload("Asset").Where("user_id = ?", uId).Find(&favorites).Error; err != nil {
		return nil, err
	}

	var favoriteAssets []models.AssetResponse
	for _, favorite := range favorites {
		var assetResponse models.AssetResponse

		result, err := helpers.AggregateAssetType(favorite.Asset)
		if err != nil {
			return nil, err
		}
		assetResponse = *result

		favoriteAssets = append(favoriteAssets, assetResponse)
	}

	return favoriteAssets, nil
}

func AddToFavorites(uId uint, aId uint) (*models.Favorite, error) {
	db := database.GetDBConnection()
	favorite := models.Favorite{UserID: uId, AssetID: aId}
	if err := db.Create(&favorite).Error; err != nil {
		return nil, fmt.Errorf("something went wrong while adding asset with id %v to favorites", aId)

	}
	return &favorite, nil
}

func RemoveFromFavorites(fId uint) (*uint, error) {
	db := database.GetDBConnection()

	if err := db.Delete(&models.Favorite{}, fId).Error; err != nil {

		return nil, fmt.Errorf("something went wrong while deleting favorite %v", fId)

	}
	return &fId, nil
}
