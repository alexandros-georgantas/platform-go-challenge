package services

import (
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/helpers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"gorm.io/gorm"
)

type FavoriteService interface {
	GetFavorites(uId uint) ([]models.AssetResponse, error)
	AddToFavorites(uId uint, aId uint) (*models.Favorite, error)
	RemoveFromFavorites(fId uint) (*uint, error)
}

type favoriteService struct {
	db gorm.DB
}

func NewFavoriteService(db gorm.DB) (FavoriteService, error) {
	return &favoriteService{db: db}, nil
}

func (fs *favoriteService) GetFavorites(uId uint) ([]models.AssetResponse, error) {
	var favorites []models.Favorite

	if err := fs.db.Preload("Asset").Where("user_id = ?", uId).Find(&favorites).Error; err != nil {
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

func (fs *favoriteService) AddToFavorites(uId uint, aId uint) (*models.Favorite, error) {
	favorite := models.Favorite{UserID: uId, AssetID: aId}

	if err := fs.db.Create(&favorite).Error; err != nil {
		return nil, fmt.Errorf("something went wrong while adding asset with id %v to favorites", aId)

	}
	return &favorite, nil
}

func (fs *favoriteService) RemoveFromFavorites(fId uint) (*uint, error) {

	if err := fs.db.Delete(&models.Favorite{}, fId).Error; err != nil {

		return nil, fmt.Errorf("something went wrong while deleting favorite %v", fId)

	}
	return &fId, nil
}
