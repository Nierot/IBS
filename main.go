package main

import (
	"fmt"

	"github.com/Nierot/InvictusBackend/controllers"
	"github.com/Nierot/InvictusBackend/middleware"
	"github.com/Nierot/InvictusBackend/models"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	makeConfig()

	discord := setupDiscord()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*.tmpl")

	gin.SetMode(viper.GetString("Server.GinMode"))

	rest := r.Group(viper.GetString("Server.Path"))

	models.SetupDB()

	/*
		Templates
	*/
	templates := r.Group("").
		Use(middleware.APITokenAuth()).
		Use(middleware.JWTAuth()).
		Use(middleware.LoginMiddleware())
	{
		templates.GET("/", controllers.IndexController)
		templates.GET("/images", controllers.ImagesController)
		templates.GET("/sales", controllers.SalesController)
		templates.GET("/purchases", controllers.PurchasesController)
		templates.GET("/purchases/new", controllers.NewPurchaseController)
		templates.GET("/users", controllers.UsersController)
		templates.GET("/music", controllers.MusicController)
		templates.GET("/settings", controllers.SettingsController)
		templates.GET("/products", controllers.ProductsController)
		templates.GET("/tally", controllers.TallyController)
		templates.GET("/statistics", controllers.StatisticsController)
	}

	/*
		Authentication related routes
	*/
	auth := rest.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/token", controllers.GenerateToken)
		auth.POST("/register", controllers.RegisterUser)
	}

	secured := rest.Group("")
	secured.Use(middleware.APITokenAuth())
	secured.Use(middleware.JWTAuth())

	secured.POST("/token/create", controllers.CreateAPIToken)

	products := secured.Group("/products")
	{
		products.GET("/", controllers.FindProducts)
		products.POST("/", controllers.CreateProduct)
		products.GET("/unique", controllers.FindUniqueProducts)
	}

	purchases := secured.Group("/purchases")
	{
		purchases.POST("/", controllers.CreatePurchase)
		purchases.GET("/all", controllers.GetAllPurchases)
		purchases.GET("/:id", controllers.GetPurchasesPerUser)
	}

	sales := secured.Group("/sales")
	{
		sales.POST("/", controllers.CreateSale)
		sales.POST("/tally", controllers.InputTallySheet)
		sales.GET("/", controllers.GetAllSales)
	}

	/*
		Image related routes
	*/
	secured.GET("/image", controllers.RandomImage)

	if viper.GetBool("Image.Scan.Enabled") {
		go controllers.ImageScanner()
	}

	/*
		Discord related routes
	*/
	quotes := secured.Group("/quotes")
	{
		quotes.GET("/random", controllers.GetRandomQuote)
		quotes.GET("/", controllers.GetAllQuotes)
	}

	if viper.GetBool("Discord.Scan.Enabled") {
		go models.QuoteScanner(discord)
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

	viper.SetDefault("Auth.TokenAge", 8600)

	viper.SetDefault("Image.Path", "./images")

	viper.SetDefault("Image.Scan.Enabled", true)
	viper.SetDefault("Image.Scan.Interval", 5)

	viper.SetDefault("Discord.Token", "<BotToken>")
	viper.SetDefault("Discord.QuoteChannel", "<QuoteChannelID>")
	viper.SetDefault("Discord.Scan.Enabled", true)
	viper.SetDefault("Discord.Scan.Interval", 5)

	viper.ReadInConfig()
	viper.SafeWriteConfig()
}

func setupDiscord() *discordgo.Session {
	discord, err := discordgo.New("Bot " + viper.GetString("Discord.Token"))

	if err != nil {
		fmt.Println("Discord token not working!")
		panic(err)
	}

	return discord
}
