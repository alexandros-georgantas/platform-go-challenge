package middlewares

import (
	"net/http"

	"github.com/alexandros-georgantas/platform-go-challenge/internal/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		c.Abort()
		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := utils.VerifyToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
