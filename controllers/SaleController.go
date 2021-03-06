package controllers

import (
	"encoding/json"
	"io/ioutil"
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
	}

	models.DB.Find(&product, "id = ?", input.ProductID)

	models.DB.Create(&s)

	s.Product = product

	c.JSON(http.StatusOK, gin.H{"Sale": s})
}

func GetAllSales(c *gin.Context) {
	var sales []models.Sale
	var products []models.Product

	salesTx := models.DB.Find(&sales)
	productsTx := models.DB.Find(&products)

	if salesTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": salesTx.Error.Error()})
		return
	}

	if productsTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": productsTx.Error.Error()})
		return
	}

	for idx, sale := range sales {
		sales[idx].Product = products[sale.ProductID-1]
	}

	c.JSON(http.StatusOK, gin.H{"Sales": sales})
}

func InputTallySheet(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tallySheet map[int]map[int]int

	if err := json.Unmarshal(data, &tallySheet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var products []models.Product
	models.DB.Find(&products).Where("depleted = ?", false)

	for user, sheet := range tallySheet {
		for product, amount := range sheet {
			if amount == 0 {
				continue
			}

			sale := models.Sale{
				UserID:    user,
				ProductID: product,
				Amount:    amount,
				Settled:   false,
			}

			models.DB.Create(&sale)
		}
	}

	c.JSON(http.StatusOK, gin.H{"TallySheet": tallySheet})
}
