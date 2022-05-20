package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserID    int     `json:"UserID"`
	ProductID int     `json:"ProductID"`
	Product   Product `json:"Product"`
	Amount    int     `json:"Amount"`
	Price     float32 `json:"Price"`
}

type SaleInput struct {
	UserID    int     `json:"UserID" binding:"required"`
	ProductID int     `json:"ProductID" binding:"required"`
	Amount    int     `json:"Amount" binding:"required"`
	Price     float32 `json:"Price" binding:"required"`
}
