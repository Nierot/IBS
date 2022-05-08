package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `json:"Name"`
	Volume  int     `json:"Content"`
	Alcohol float32 `json:"Alcohol"`
}

type ProductInput struct {
	Name    string  `json:"Name" binding:"required"`
	Volume  int     `json:"Content" binding:"required"`
	Alcohol float32 `json:"Alcohol" binding:"required"`
}
