package services

import (
	"errors"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/database"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/serializers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/utils"
	"gorm.io/gorm"
)

// var (
// 	db = database.GetDBConnection()
// )

func SignUp(su *serializers.SignUpUser) (*uint, error) {
	db := database.GetDBConnection()
	user := models.User{GivenName: su.GivenName, Surname: su.Surname, Email: su.Email, Password: su.Password}
	dbErr := db.Create(&user).Error

	if dbErr != nil {
		return nil, errors.New("could not create user")
	}

	return &user.ID, nil
}

func Login(uc *serializers.UserCredentials) (*string, error) {
	db := database.GetDBConnection()
	user := models.User{}

	dbErr := db.Where("email = ?", uc.Email).First(&user).Error

	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		return nil, errors.New("wrong user credentials")
	}

	if !utils.VerifyPassword(uc.Password, user.Password) {
		return nil, errors.New("wrong user credentials")
	}

	token, tErr := utils.CreateToken(user.ID)

	if tErr != nil {
		return nil, tErr
	}

	return token, nil
}
