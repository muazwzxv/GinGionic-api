package models

import (
	"Go-Learn-API/auth"
)

// Auth := Auth todo
type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

// FetchAuth := Fetch
func (s *Server) FetchAuth(auth *auth.AuthDetails) (*Auth, error) {
	authenticate := &Auth{}
	err := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", auth.UserID, auth.AuthUUID).Take(&authenticate).Error

	if err != nil {
		return nil, err
	}
	return authenticate, nil
}

// DeleteAuth := delete auth info
func (s *Server) DeleteAuth(auth *auth.AuthDetails) error {
	authenticate := &Auth{}
	db := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", auth.UserID, auth.AuthUUID).Take(&authenticate).Delete(&authenticate)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
