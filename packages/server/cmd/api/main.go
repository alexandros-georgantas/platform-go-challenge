package main

import (
	"fmt"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/database"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/models"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/server"
)

func init() {
	db := database.GetDBConnection()
	db.AutoMigrate(&models.User{}, &models.Audience{}, &models.Chart{}, &models.Insight{}, &models.Asset{}, &models.Favorite{})

}

func main() {

	server := server.Create()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
