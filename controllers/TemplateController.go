package controllers

import (
	"fmt"
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
	var sales []models.SaleJoin

	models.DB.
		Model(&models.Sale{}).
		Select(`
			sales.amount as sales_amount, 
			sales.settled as settled, 
			sales.created_at as sale_created_at,
			users.username as username, 
			products.name as product_name, 
			products.volume as volume, 
			products.alcohol as alcohol`).
		Joins("INNER JOIN users ON sales.user_id = users.id").
		Joins("INNER JOIN products ON sales.product_id = products.id").
		Scan(&sales)

	fmt.Println(sales)

	c.HTML(http.StatusOK, "sales.tmpl", gin.H{
		"title": "Verkoop",
		"api":   viper.GetString("Server.Path"),
		"sales": sales,
	})
}

func PurchasesController(c *gin.Context) {
	var (
		purchases []models.Purchase
		products  []models.Product
	)

	models.DB.Find(&purchases)
	models.DB.Find(&products)

	c.HTML(http.StatusOK, "purchases.tmpl", gin.H{
		"title":     "Inkoop",
		"api":       viper.GetString("Server.Path"),
		"purchases": purchases,
		"products":  products,
	})
}

func NewPurchaseController(c *gin.Context) {
	var (
		products []models.Product
		users    []models.User
	)

	models.DB.Find(&products)
	models.DB.Model(&models.User{}).
		Select("id, name, username").
		Scan(&users)

	c.HTML(http.StatusOK, "new_purchase.tmpl", gin.H{
		"title":    "Nieuwe inkoop",
		"api":      viper.GetString("Server.Path"),
		"products": products,
		"users":    users,
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

func MusicController(c *gin.Context) {
	c.HTML(http.StatusOK, "music.tmpl", gin.H{
		"title": "Muziek",
		"api":   viper.GetString("Server.Path"),
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

func TallyController(c *gin.Context) {
	var (
		users     []models.User
		products  []models.Product
		purchases []models.Purchase
	)

	models.DB.Find(&products)
	models.DB.Find(&purchases)
	models.DB.Model(&models.User{}).
		Select("id, name, username").
		Scan(&users)

	c.HTML(http.StatusOK, "tally.tmpl", gin.H{
		"title":     "Verwerk een streeplijst",
		"api":       viper.GetString("Server.Path"),
		"users":     users,
		"products":  products,
		"purchases": purchases,
	})
}

func StatisticsController(c *gin.Context) {
	c.HTML(http.StatusOK, "statistics.tmpl", gin.H{
		"title": "Instellingen",
		"api":   viper.GetString("Server.Path"),
	})
}
