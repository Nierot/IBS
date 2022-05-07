package main

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/controllers"
	"github.com/Nierot/InvictusBackend/middleware"
	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	rest := r.Group("/api")

	models.SetupDB()
	{
		rest.POST("/token", controllers.GenerateToken)
		rest.POST("/user/register", controllers.RegisterUser)
	}

	products := rest.Group("/products")
	{
		products.GET("/", controllers.FindProducts)
		products.GET("/unique", controllers.FindUniqueProducts)
		products.POST("/", controllers.CreateProduct)
	}

	a := rest.Group("/auth")
	{
		a.POST("/token", controllers.GenerateToken)
		a.POST("/register", controllers.RegisterUser)
	}

	secured := rest.Group("/secured").Use(middleware.Auth())
	{
		secured.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	r.Run()
}
