package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("ENCRYPTION_KEY")
	if requiredToken == "" {
		log.Fatal("API_TOKEN environment variable not set")
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
		}
		if token != requiredToken {
			c.JSON(401, gin.H{"error": "Invalid token"})
		}
		c.Next()
	}

}
