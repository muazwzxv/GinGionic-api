package models

import (
	"Go-Learn-API/auth"
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
	ConnectDB() (*gorm.DB, error)

	// user methods
	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

	// Todo methods:
	CreateTodo(*Todo) (*Todo, error)

	// Book method
	CreateBooks(*Book) (*Book, error)
	GetAllBooks() ([]Book, error)
	GetByIDBooks(int) (Book, error)

	// auth methods
	FetchAuth(*auth.AuthDetails) (*Auth, error)
	DeleteAuth(*auth.AuthDetails) error
	CreateAuth(uint64) (*Auth, error)
}

var (
	// Model := the instance spawned
	Model ServerInterface = &Server{}
)

// ConnectDB connects to the database
func (s *Server) ConnectDB() (*gorm.DB, error) {

	dbConfig, err := configs.DBConfig()

	if err != nil {
		panic("Failed to fetch database")
	}

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

	s.DB.Debug().AutoMigrate(
		&Book{},
		&User{},
	)

	return s.DB, nil
}
