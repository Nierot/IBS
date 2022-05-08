package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
)

type APITokenRequest struct {
	Device string `json:"device" binding:"required"`
}

func CreateAPIToken(c *gin.Context) {
	var req APITokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	token := models.GenerateToken(req.Device)
	models.DB.Create(&token)

	c.JSON(http.StatusCreated, gin.H{"Token": token})
}
