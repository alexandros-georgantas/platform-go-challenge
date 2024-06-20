package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func GetAssets(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))

	assets, err := services.GetAssets(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"assets": assets,
	})
}

func GetCharts(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))

	charts, err := services.GetCharts(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"assets": charts,
	})
}
func GetAudiences(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))

	audiences, err := services.GetAudiences(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"assets": audiences,
	})
}

func GetInsights(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))

	insights, err := services.GetInsights(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"assets": insights,
	})
}

func UpdateDescription(c *gin.Context) {
	var uD serializers.UpdateDescription
	aId, _ := strconv.Atoi(c.Param("assetId"))
	bErr := c.BindJSON(&uD)
	uA, aErr := services.UpdateDescription(uint(aId), uD.Description)

	if uD.Description == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid request").Error(),
		})
		return
	}

	if aErr != nil || bErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("something went wrong while updating description of asset with id %v", aId).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"asset": uA,
	})

}
