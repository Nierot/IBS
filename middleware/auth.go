package middleware

import (
	"net/http"
	"time"

	"github.com/Nierot/InvictusBackend/auth"
	"github.com/Nierot/InvictusBackend/models"
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
		a, err := c.Cookie("Authorization")

		if tokenString == "" && err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}

		if tokenString == "" {
			tokenString = a
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("authorized", true)
		c.Next()
	}
}

func APITokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		var apiToken models.APIToken

		tx := models.DB.Limit(1).Find(&apiToken, "token = ?", tokenString)

		if tx.Error != nil {
			c.Set("api_token", false)
		} else if apiToken.Token == tokenString {
			c.Set("api_token", true)
			c.Set("authorized", true)
			tx.Update("last_used", time.Now())
		}

		c.Next()
	}
}
