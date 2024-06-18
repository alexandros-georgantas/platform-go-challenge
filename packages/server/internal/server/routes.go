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

	// API ROUTES
	v1 := r.Group("/api/v1")

	{
		v1.POST("/users", controllers.SignUp)
		v1.POST("/tokens", controllers.Login)
	}

	v1.Use(middlewares.Authenticate)

	{
		v1.GET("/users/:userId/favorites", s.healthHandler)    //get user's favorites
		v1.POST("/users/:userId/favorites", s.healthHandler)   //set user's favorite
		v1.DELETE("/users/:userId/favorites", s.healthHandler) //delete user's favorite
		v1.GET("/assets", s.healthHandler)                     //get all assets
		v1.GET("/assets/charts", s.healthHandler)              //get charts
		v1.GET("/assets/audiences", s.healthHandler)           //get audiences
		v1.GET("/assets/insights", s.healthHandler)            //get insights
	}

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
