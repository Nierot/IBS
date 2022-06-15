package controllers

import (
	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

func GetRandomQuote(c *gin.Context) {
	c.JSON(501, gin.H{
		"Error": "Not implemented yet!",
	})
}

func GetAllQuotes(c *gin.Context) {
	var quotes []models.Quote
	var quotesJSON []models.QuoteJSON

	tx := models.DB.Find(&quotes)

	for _, quote := range quotes {
		quotesJSON = append(quotesJSON, models.QuoteJSON{
			Quote:       quote.Quote,
			AddedBy:     quote.AddedBy,
			Timestamp:   quote.Timestamp,
			MessageID:   quote.MessageID,
			Attachments: models.StringToSlice(quote.Attachments),
			Reactions:   models.StringToSlice(quote.Reactions),
		})
	}

	if tx.Error != nil {
		c.JSON(500, gin.H{
			"Error": tx.Error,
		})
		return
	}

	c.JSON(200, quotesJSON)
}
