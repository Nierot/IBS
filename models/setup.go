package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&APIToken{})
	db.AutoMigrate(&Sale{})
	db.AutoMigrate(&Purchase{})
	db.AutoMigrate(&Quote{})

	DB = db
}
