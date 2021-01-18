package models

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
)

// User := user properties
type User struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
	//DeletedAt gorm.DeletedAt `gorm: "index"`
	CreatedAt time.Time
}

// ValidateEmail := validate incoming email
func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.New("Invalid email")
	}
	return nil
}

// CreateUser := create user
func (s *Server) CreateUser(user *User) (*User, error) {

	emailError := s.ValidateEmail(user.Email)
	if emailError != nil {
		return nil, emailError
	}

	err := s.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail := returns user
func (s *Server) GetUserByEmail(email string) (*User, error) {

	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
