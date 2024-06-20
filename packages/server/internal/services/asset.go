package services

import (
	"errors"
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/helpers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"gorm.io/gorm"
)

type AssetService interface {
	GetAsset(aid uint) (*models.AssetResponse, error)
	GetAssets(p int, pS int) ([]models.AssetResponse, error)
	GetCharts(p int, pS int) ([]models.Chart, error)
	GetAudiences(p int, pS int) ([]models.Audience, error)
	GetInsights(p int, pS int) ([]models.Insight, error)
	UpdateDescription(aId uint, d string) (*models.AssetResponse, error)
}

type assetService struct {
	db gorm.DB
}

func NewAssetService(db gorm.DB) (AssetService, error) {
	return &assetService{db: db}, nil
}

func (as *assetService) GetAssets(p int, pS int) ([]models.AssetResponse, error) {
	var assets []models.Asset
	var assetResponses []models.AssetResponse

	if err := as.db.Scopes(helpers.Paginate(p, pS)).Find(&assets).Error; err != nil {
		return nil, errors.New("something went wrong while fetching assets")
	}

	for _, asset := range assets {
		var assetResponse models.AssetResponse
		result, err := helpers.AggregateAssetType(asset, as.db)

		if err != nil {
			return nil, err
		}

		assetResponse = *result

		assetResponses = append(assetResponses, assetResponse)
	}
	return assetResponses, nil
}

func (as *assetService) GetCharts(p int, pS int) ([]models.Chart, error) {
	var charts []models.Chart

	if err := as.db.Preload("Asset").Scopes(helpers.Paginate(p, pS)).Find(&charts).Error; err != nil {
		return nil, errors.New("something went wrong while fetching charts")

	}
	return charts, nil
}

func (as *assetService) GetAudiences(p int, pS int) ([]models.Audience, error) {
	var audiences []models.Audience

	if err := as.db.Preload("Asset").Scopes(helpers.Paginate(p, pS)).Find(&audiences).Error; err != nil {
		return nil, errors.New("something went wrong while fetching audiences")

	}

	return audiences, nil
}

func (as *assetService) GetInsights(p int, pS int) ([]models.Insight, error) {
	var insights []models.Insight

	if err := as.db.Preload("Asset").Scopes(helpers.Paginate(p, pS)).Find(&insights).Error; err != nil {
		return nil, errors.New("something went wrong while fetching insights")

	}
	return insights, nil
}

func (as *assetService) UpdateDescription(aId uint, d string) (*models.AssetResponse, error) {
	var asset models.Asset

	if fErr := as.db.First(&asset, aId).Error; fErr != nil {
		return nil, fmt.Errorf("something went wrong while fetching  asset with id %v", aId)
	}

	if pErr := as.db.Model(&asset).Update("description", d).Error; pErr != nil {
		return nil, fmt.Errorf("something went wrong while updating description of asset with id %v", aId)
	}

	result, aErr := helpers.AggregateAssetType(asset, as.db)

	if aErr != nil {
		return nil, aErr
	}

	return result, nil
}

func (as *assetService) GetAsset(aid uint) (*models.AssetResponse, error) {
	var asset models.Asset
	if err := as.db.First(&asset, aid).Error; err != nil {
		return nil, fmt.Errorf("something went wrong while fetching asset with id %v", aid)
	}
	result, aErr := helpers.AggregateAssetType(asset, as.db)

	if aErr != nil {
		return nil, aErr
	}

	return result, nil
}
