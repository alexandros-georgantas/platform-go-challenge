package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/health", s.healthHandler)

	// v1 := r.Group("/api/v1")

	return r
}

func (s *Server) healthHandler(c *gin.Context) {

	resp := make(map[string]string)
	resp["msg"] = "OK"
	c.JSON(http.StatusOK, resp)
}
