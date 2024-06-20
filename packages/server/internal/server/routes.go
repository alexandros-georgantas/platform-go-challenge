package server

import (
	"net/http"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/controllers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/middlewares"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// SERVICES
	userService, _ := services.NewUserService(*s.db)
	assetService, _ := services.NewAssetService(*s.db)
	favoriteService, _ := services.NewFavoriteService(*s.db)

	// CONTROLLERS
	userController, _ := controllers.NewUserController(userService)
	assetController, _ := controllers.NewAssetController(assetService)
	favoriteController, _ := controllers.NewFavoriteController(favoriteService)

	// GENERIC ROUTE
	r.GET("/health", s.healthHandler)

	// API OPEN ROUTES
	v1 := r.Group("/api/v1")

	{
		v1.POST("/users", userController.SignUp)
		v1.POST("/tokens", userController.Login)
	}

	v1.Use(middlewares.Authenticate)

	// API PROTECTED ROUTES
	{
		v1.GET("/assets", assetController.GetAssets) //get all assets
		// v1.GET("/assets/:assetId", assetController.UpdateDescription)               //patch asset's description
		v1.PATCH("/assets/:assetId", assetController.UpdateDescription)                      //patch asset's description
		v1.GET("/assets/charts", assetController.GetCharts)                                  //get charts
		v1.GET("/assets/insights", assetController.GetInsights)                              //get insights
		v1.GET("/assets/audiences", assetController.GetAudiences)                            //get audiences
		v1.GET("/users/:userId/favorites", favoriteController.GetFavorites)                  //get user's favorites
		v1.POST("/users/:userId/favorites", favoriteController.AddToFavorites)               //set user's favorite
		v1.DELETE("/users/:userId/favorites/:favoriteId", favoriteController.DeleteFavorite) //delete user's favorite

	}

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
