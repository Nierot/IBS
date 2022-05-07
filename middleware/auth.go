package middleware

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/auth"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request has a valid API token, so we do not check the JWT
		if c.GetBool("api_token") {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func APITokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		c.Set("api_token", tokenString == "aaa")

		c.Next()
	}
}
