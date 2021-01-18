package models

import (
	"errors"

	"github.com/badoux/checkmail"
)

// User := user properties
type User struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
}

// ValidateEmail := validate incoming email
func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	} else {
		if err := checkmail.ValidateFormat(email); err != nil {
			return errors.New("Invalid email")
		}
	}
	return nil
}
