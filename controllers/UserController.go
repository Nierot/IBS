package controllers

import (
	"net/http"

	"github.com/Nierot/InvictusBackend/auth"
	"github.com/Nierot/InvictusBackend/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type TokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	record := models.DB.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": record.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ID": user.ID, "Username": user.Username, "UserEmail": user.Email})
}

type Form struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var (
		f    Form
		user models.User
	)

	if err := c.Bind(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	record := models.DB.Where("email = ?", f.Email).First(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": record.Error.Error()})
		c.Abort()
		return
	}

	if !user.CheckPassword(f.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   viper.GetInt("Auth.TokenAge"),
		Secure:   viper.GetString("Server.GinMode") != "debug",
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.Redirect(http.StatusFound, "/")
}

func GenerateToken(c *gin.Context) {
	var (
		req  TokenRequest
		user models.User
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	record := models.DB.Where("email = ?", req.Email).First(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": record.Error.Error()})
		c.Abort()
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": tokenString})
}
