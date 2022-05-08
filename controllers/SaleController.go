package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func CreateSale(c *gin.Context) {
	var input models.SaleInput
	var product models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Sale{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Amount:    input.Amount,
		Price:     input.Price,
	}

	models.DB.Find(&product, "id = ?", input.ProductID)

	models.DB.Create(&s)

	s.Product = product

	c.JSON(http.StatusOK, gin.H{"Sale": s})
}
