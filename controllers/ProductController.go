package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func FindProducts(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"Products": products})
}

func FindUniqueProducts(c *gin.Context) {
	var products []models.Product
	var unique []string

	u := make(map[string]bool)

	models.DB.Find(&products)

	for p := range products {
		name := products[p].Name

		if u[name] {
			continue
		}

		u[name] = true
		unique = append(unique, name)
	}

	c.JSON(http.StatusOK, gin.H{"Products": unique})
}

func CreateProduct(c *gin.Context) {
	var input models.ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alcohol, aErr := input.Alcohol.Float64()
	volume, vErr := input.Volume.Float64()

	if aErr != nil || vErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": gin.H{"Alcohol": aErr, "Volume": vErr}})
		return
	}

	p := models.Product{Name: input.Name, Alcohol: alcohol, Volume: volume}
	models.DB.Create(&p)

	c.JSON(http.StatusOK, gin.H{"Product": p})
}
