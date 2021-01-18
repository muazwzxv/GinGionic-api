package models

// User := user properties
type User struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
}
