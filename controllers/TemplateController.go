package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Index",
		"api":   viper.GetString("Server.Path"),
	})
}

func ImagesController(c *gin.Context) {
	c.HTML(http.StatusOK, "images.tmpl", gin.H{
		"title": "Fotos",
		"api":   viper.GetString("Server.Path"),
	})
}

func SalesController(c *gin.Context) {
	c.HTML(http.StatusOK, "sales.tmpl", gin.H{
		"title": "Verkoop",
		"api":   viper.GetString("Server.Path"),
	})
}

func PurchasesController(c *gin.Context) {
	c.HTML(http.StatusOK, "purchases.tmpl", gin.H{
		"title": "Inkoop",
		"api":   viper.GetString("Server.Path"),
	})
}

func UsersController(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.HTML(http.StatusOK, "users.tmpl", gin.H{
		"title": "Gebruikers",
		"api":   viper.GetString("Server.Path"),
		"users": users,
	})
}

func SettingsController(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.tmpl", gin.H{
		"title": "Instellingen",
		"api":   viper.GetString("Server.Path"),
	})
}

func ProductsController(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)

	c.HTML(http.StatusOK, "products.tmpl", gin.H{
		"title":    "Producten",
		"api":      viper.GetString("Server.Path"),
		"products": products,
	})
}
