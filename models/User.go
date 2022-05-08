package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"Name" binding:"required"`
	Username string `json:"Username" gorm:"unique" binding:"required"`
	Email    string `json:"Email" gorm:"unique" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

// https://codewithmukesh.com/blog/jwt-authentication-in-golang/

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
