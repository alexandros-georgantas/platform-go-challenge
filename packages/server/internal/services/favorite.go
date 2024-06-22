package services

import (
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/helpers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"gorm.io/gorm"
)

type FavoriteService interface {
	GetFavorite(uId uint, aId uint) (*models.AssetResponse, error)
	GetFavorites(uId uint) (serializers.FavoritesResponse, error)
	AddToFavorites(uId uint, aId uint) (*models.Favorite, error)
	RemoveFromFavorites(fId uint) (*uint, error)
}

type favoriteService struct {
	db gorm.DB
}

func NewFavoriteService(db gorm.DB) (FavoriteService, error) {
	return &favoriteService{db: db}, nil
}

func (fs *favoriteService) GetFavorite(uId uint, aId uint) (*models.AssetResponse, error) {
	var favorite models.Favorite

	if err := fs.db.Joins("Asset").Where("user_id = ?", uId).First(&favorite, aId).Error; err != nil {
		return nil, err
	}

	result, aErr := helpers.AggregateAssetType(favorite.Asset, fs.db)

	if aErr != nil {
		return nil, aErr
	}

	return result, nil
}

func (fs *favoriteService) GetFavorites(uId uint) (serializers.FavoritesResponse, error) {
	var favorites []models.Favorite
	pr := serializers.FavoritesResponse{}
	var count int64
	pr.Favorites = nil
	pr.TotalCount = 0

	fs.db.Model(&models.Favorite{}).Where("user_id = ?", uId).Count(&count)

	if err := fs.db.Joins("Asset").Where("user_id = ?", uId).Find(&favorites).Error; err != nil {
		return pr, err
	}

	var favoriteAssets []models.AssetResponse
	for _, favorite := range favorites {
		var assetResponse models.AssetResponse

		result, err := helpers.AggregateAssetType(favorite.Asset, fs.db)
		if err != nil {
			return pr, err
		}
		assetResponse = *result

		favoriteAssets = append(favoriteAssets, assetResponse)
	}

	pr.Favorites = &favoriteAssets
	pr.TotalCount = int(count)

	return pr, nil
}

func (fs *favoriteService) AddToFavorites(uId uint, aId uint) (*models.Favorite, error) {
	favorite := models.Favorite{UserID: uId, AssetID: aId}
	fmt.Println("here")
	if err := fs.db.Preload("Asset").Create(&favorite).First(&favorite).Error; err != nil {
		return nil, fmt.Errorf("something went wrong while adding asset with id %v to favorites", aId)

	}
	return &favorite, nil
}

func (fs *favoriteService) RemoveFromFavorites(fId uint) (*uint, error) {
	fmt.Println("edo")
	if err := fs.db.Delete(&models.Favorite{}, fId).Error; err != nil {

		return nil, fmt.Errorf("something went wrong while deleting favorite %v", fId)

	}

	return &fId, nil
}
