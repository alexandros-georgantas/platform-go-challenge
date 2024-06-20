package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func AddToFavorites(c *gin.Context) {
	var aF serializers.AddToFavorites
	pUId, _ := strconv.Atoi(c.Param("userId"))
	cUId := c.MustGet("userId")
	bErr := c.BindJSON(&aF)

	if pUId != cUId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token action"})
		return
	}

	favorite, aErr := services.AddToFavorites(uint(pUId), uint(aF.ID))

	if aErr != nil || bErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("something went wrong while adding asset with id %v to favorites", aF.ID).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"assets": favorite,
	})
}

func DeleteFavorite(c *gin.Context) {
	pUId, _ := strconv.Atoi(c.Param("userId"))
	pFId, _ := strconv.Atoi(c.Param("favorite"))
	cUId := c.MustGet("userId")

	if pUId != cUId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token action"})
		return
	}

	dId, aErr := services.RemoveFromFavorites(uint(pFId))

	if aErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("something went wrong while removing asset with id %v from favorites", pFId).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"asset": dId,
	})
}
