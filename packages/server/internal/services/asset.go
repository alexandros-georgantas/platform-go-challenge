package services

import (
	"errors"
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/database"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/helpers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/utils"
)

func GetAssets(p int, pS int) ([]models.AssetResponse, error) {
	db := database.GetDBConnection()
	var assets []models.Asset
	var assetResponses []models.AssetResponse

	if err := db.Scopes(utils.Paginate(p, pS)).Find(&assets).Error; err != nil {
		return nil, errors.New("something went wrong while fetching assets")
	}

	for _, asset := range assets {
		var assetResponse models.AssetResponse
		result, err := helpers.AggregateAssetType(asset)

		if err != nil {
			return nil, err
		}

		assetResponse = *result

		assetResponses = append(assetResponses, assetResponse)
	}
	return assetResponses, nil
}

func GetCharts(p int, pS int) ([]models.Chart, error) {
	db := database.GetDBConnection()
	var charts []models.Chart

	if err := db.Preload("Asset").Scopes(utils.Paginate(p, pS)).Find(&charts).Error; err != nil {
		return nil, errors.New("something went wrong while fetching charts")

	}
	return charts, nil
}

func GetAudiences(p int, pS int) ([]models.Audience, error) {
	db := database.GetDBConnection()
	var audiences []models.Audience

	if err := db.Preload("Asset").Scopes(utils.Paginate(p, pS)).Find(&audiences).Error; err != nil {
		return nil, errors.New("something went wrong while fetching audiences")

	}

	return audiences, nil
}

func GetInsights(p int, pS int) ([]models.Insight, error) {
	db := database.GetDBConnection()
	var insights []models.Insight

	if err := db.Preload("Asset").Scopes(utils.Paginate(p, pS)).Find(&insights).Error; err != nil {
		return nil, errors.New("something went wrong while fetching insights")

	}
	return insights, nil
}

func UpdateDescription(aId uint, d string) (*models.AssetResponse, error) {
	db := database.GetDBConnection()
	var asset models.Asset

	if fErr := db.First(&asset, aId).Error; fErr != nil {
		return nil, fmt.Errorf("something went wrong while fetching  asset with id %v", aId)
	}

	if pErr := db.Model(&asset).Update("description", d).Error; pErr != nil {
		return nil, fmt.Errorf("something went wrong while updating description of asset with id %v", aId)
	}

	result, aErr := helpers.AggregateAssetType(asset)

	if aErr != nil {
		return nil, aErr
	}

	return result, nil
}
