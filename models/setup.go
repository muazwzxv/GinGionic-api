package models

import (
	"Go-Learn-API/configs"
	"fmt"

	"github.com/jinzhu/gorm"

	// mysql dialect in gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Server struct
type Server struct {
	DB *gorm.DB
}

// ServerInterface := server interface
type ServerInterface interface {
	ConnectDB()
}

// ConnectDB connects to the database
func (s *Server) ConnectDB() (*gorm.DB, error) {

	dbConfig, err := configs.DBConfig()

	if err != nil {
		panic("Failed to fetch database")
	}

	// db, err := gorm.Open("mysql", "user:password@(localhost:3306)/gormLearn?charset=utf8&parseTime=True&loc=Local")
	s.DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
		),
	)

	if err != nil {
		panic("Failed to connect to database!")
	}

	s.DB.Debug().AutoMigrate(&Book{})

	return s.DB, nil
}
