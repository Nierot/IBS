package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	var input models.PurchaseInput
	var product models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Purchase{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Price:     input.Price,
		Deposit:   input.Deposit,
		Amount:    input.Amount,
	}

	models.DB.First(&product, "id = ?", input.ProductID)

	models.DB.Create(&p)

	p.Product = product

	c.JSON(http.StatusOK, gin.H{"Purchase": p})
}

func GetAllPurchases(c *gin.Context) {
	var purchases []models.Purchase
	var products []models.Product

	purchasesTx := models.DB.Find(&purchases)
	productsTx := models.DB.Find(&products)

	if purchasesTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": purchasesTx.Error.Error()})
		return
	}

	if productsTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": productsTx.Error.Error()})
		return
	}

	for idx, purchase := range purchases {
		purchases[idx].Product = products[purchase.ProductID-1]
	}

	c.JSON(http.StatusOK, gin.H{"Purchases": purchases})
}

func GetPurchasesPerUser(c *gin.Context) {
	var purchases []models.Purchase
	var products []models.Product

	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	purchasesTx := models.DB.Where("user_id = ?", userID).Find(&purchases)
	productsTx := models.DB.Find(&products)

	if purchasesTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": purchasesTx.Error.Error()})
		return
	}

	if productsTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": productsTx.Error.Error()})
		return
	}

	for idx, purchase := range purchases {
		purchases[idx].Product = products[purchase.ProductID-1]
	}

	c.JSON(http.StatusOK, gin.H{"Purchases": purchases})
}
