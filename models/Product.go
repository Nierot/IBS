package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Content int     `json:"content"`
	Amount  int     `json:"amount"`
}

type ProductInput struct {
	Name    string  `json:"name" binding:"required"`
	Price   float32 `json:"price" binding:"required"`
	Content int     `json:"content" binding:"required"`
	Amount  int     `json:"amount" binding:"required"`
}
