package models

import "errors"

// Book := Model for book
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateBooks := create book row
func (s *Server) CreateBooks(book *Book) (*Book, error) {
	if book.Title == "" {
		return nil, errors.New("please provide valid title")
	}

	if book.Author == "" {
		return nil, errors.New("A valid author is required")
	}

	err := s.DB.Debug().Create(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}