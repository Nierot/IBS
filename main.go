package main

import (
	"github.com/Nierot/InvictusBackend/controllers"
	"github.com/Nierot/InvictusBackend/middleware"
	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	makeConfig()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors.Default())
	gin.SetMode(viper.GetString("Server.GinMode"))

	rest := r.Group(viper.GetString("Server.Path"))

	models.SetupDB()

	/*
		Authentication related routes
	*/
	auth := rest.Group("/auth")
	{
		auth.POST("/login", controllers.GenerateToken)
		auth.POST("/register", controllers.RegisterUser)
	}

	secured := rest.Group("")
	secured.Use(middleware.APITokenAuth())
	secured.Use(middleware.JWTAuth())

	products := secured.Group("/products")
	{
		products.GET("/", controllers.FindProducts)
		products.POST("/", controllers.CreateProduct)
		products.GET("/unique", controllers.FindUniqueProducts)
	}

	secured.POST("/token/create", controllers.CreateAPIToken)

	purchases := secured.Group("/purchases")
	{
		purchases.POST("/", controllers.CreatePurchase)
		purchases.GET("/all", controllers.GetAllPurchases)
	}

	/*
		Image related routes
	*/
	secured.GET("/image", controllers.RandomImage)

	if viper.GetBool("Image.Scan.Enabled") {
		go controllers.ImageScanner()
	}

	r.Run("0.0.0.0:" + viper.GetString("Server.Port"))
}

func makeConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("Server.Port", 8080)
	viper.SetDefault("Server.GinMode", "debug")
	viper.SetDefault("Server.Path", "/api")

	viper.SetDefault("Image.Path", "./images")

	viper.SetDefault("Image.Scan.Enabled", true)
	viper.SetDefault("Image.Scan.Interval", 5)

	viper.ReadInConfig()
	viper.SafeWriteConfig()
}
