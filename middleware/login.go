package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Authorization")

		if err != nil || cookie == "" {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"title": "Login",
				"api":   viper.GetString("Server.Path"),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
