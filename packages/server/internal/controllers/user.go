package controllers

import (
	"net/http"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var su serializers.SignUpUser
	err := c.BindJSON(&su)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "your request is wrong",
		})
		return
	}

	newUserId, err := services.SignUp(&su)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"userId": newUserId,
	})
}

func Login(c *gin.Context) {
	var uc serializers.UserCredentials
	bErr := c.BindJSON(&uc)

	if bErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "client request is wrong",
		})
		return
	}

	token, sErr := services.Login(&uc)

	if sErr != nil {
		if sErr.Error() == "wrong user credentials" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": sErr.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": sErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
