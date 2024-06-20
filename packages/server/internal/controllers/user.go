package controllers

import (
	"net/http"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) (UserController, error) {
	return &userController{userService: userService}, nil
}

func (uc *userController) SignUp(c *gin.Context) {
	var su serializers.SignUpUser
	err := c.BindJSON(&su)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "your request is wrong",
		})
		return
	}

	newUserId, err := uc.userService.SignUp(&su)

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

func (uc *userController) Login(c *gin.Context) {
	var sUc serializers.UserCredentials
	bErr := c.BindJSON(&uc)

	if bErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "client request is wrong",
		})
		return
	}

	token, sErr := uc.userService.Login(&sUc)

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
