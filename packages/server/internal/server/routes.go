package server

import (
	"net/http"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/controllers"
	"github.com/alexandros-georgantas/platform-go-challenge/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// GENERIC ROUTE
	r.GET("/health", s.healthHandler)

	// API OPEN ROUTES
	v1 := r.Group("/api/v1")

	{
		v1.POST("/users", controllers.SignUp)
		v1.POST("/tokens", controllers.Login)
	}

	v1.Use(middlewares.Authenticate)

	// API PROTECTED ROUTES
	{
		v1.GET("/assets", controllers.GetAssets)                    //get all assets
		v1.PATCH("/assets/:assetId", controllers.UpdateDescription) //patch asset's description
		v1.GET("/assets/charts", controllers.GetCharts)             //get charts
		v1.GET("/assets/insights", controllers.GetInsights)         //get insights
		v1.GET("/assets/audiences", controllers.GetAudiences)       //get audiences
		// v1.GET("/users/:userId/favorites", s.healthHandler)                //get user's favorites
		v1.POST("/users/:userId/favorites", controllers.AddToFavorites)               //set user's favorite
		v1.DELETE("/users/:userId/favorites/:favoriteId", controllers.DeleteFavorite) //delete user's favorite

	}

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
