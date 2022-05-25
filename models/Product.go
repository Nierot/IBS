package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name    string  `json:"Name"`
	Volume  float64 `json:"Volume"`
	Alcohol float64 `json:"Alcohol"`
}

type ProductInput struct {
	Name    string      `json:"Name" binding:"required"`
	Volume  json.Number `json:"Volume" binding:"required"`
	Alcohol json.Number `json:"Alcohol" binding:"required"`
}
