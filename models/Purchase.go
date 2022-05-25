package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	UserID     int     `json:"UserID"`
	Product    Product `json:"Product"`
	ProductID  int     `json:"ProductID"`
	Price      float32 `json:"Price"`
	Deposit    float32 `json:"Deposit"`
	Amount     int     `json:"Amount"`
	AmountSold int     `json:"AmountSold"`
	Depleted   bool    `json:"Depleted"` // when AmountSold == Amount or the bottle is empty
	Settled    bool    `json:"Settled"`  // When the purchase is settled with the originial purchaser
}

type PurchaseProductJoin struct {
	Purchase
	Product
}

type PurchaseInput struct {
	UserID    int     `json:"UserID" binding:"required"`
	ProductID int     `json:"ProductID" binding:"required"`
	Price     float32 `json:"Price" binding:"required"`
	Deposit   float32 `json:"Deposit" binding:"required"`
	Amount    int     `json:"Amount" binding:"required"`
}
