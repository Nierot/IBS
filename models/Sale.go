package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	UserID    int     `json:"UserID"`
	ProductID int     `json:"ProductID"`
	Product   Product `json:"Product"`
	Amount    int     `json:"Amount"`
	Settled   bool    `json:"Settled"`
}

type SaleInput struct {
	UserID    int `json:"UserID" binding:"required"`
	ProductID int `json:"ProductID" binding:"required"`
	Amount    int `json:"Amount" binding:"required"`
}

type SaleJoin struct {
	SalesAmount   int       `gorm:"column:sales_amount"`
	Settled       bool      `gorm:"column:settled"`
	Username      string    `gorm:"column:username"`
	ProductName   string    `gorm:"column:product_name"`
	Volume        int       `gorm:"column:volume"`
	Alcohol       float32   `gorm:"column:alcohol"`
	SaleCreatedAt time.Time `gorm:"column:sale_created_at"`
}

func SettleSales() {
	var (
		sales     []Sale
		users     []User
		purchases []Purchase
	)

	DB.Find(&sales)
	DB.Find(&users)
	DB.Find(&purchases)
	for _, sale := range sales {
		fmt.Println(sale)
	}

}
