package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("mysql", "user:password@(localhost:3306)/gormLearn?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})
	defer db.Close()

	DB = db
}
