package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	var input models.PurchaseInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// p := models.Purchase{UserID: input.UserID, ProductID: input.ProductID, Price: input.Price, Deposit: input.Deposit, Amount: input.Amount}

}
