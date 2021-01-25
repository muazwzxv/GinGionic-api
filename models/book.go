package models

import (
	"errors"
)

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

// GetAllBooks := Returns all books
func (s *Server) GetAllBooks() ([]Book, error) {
	var books []Book

	if err := s.DB.Debug().Find(&books).Error; err != nil {
		return []Book{}, errors.New("Something wrong happened")
	}

	return books, nil
}

// GetByIDBooks := Returns book by id
func (s *Server) GetByIDBooks(id int) (Book, error){
	var book Book
	if err := s.DB.Debug().Where("id = ?", id).First(&book).Error; err != nil {
		return Book{}, errors.New("Not found")
	}

	return book, nil
}

// UpdateByIDBooks := Update by id
func (s *Server) UpdateByIDBooks(id int, update *Book) (Book, error) {
	var book Book
	if err := s.DB.Debug().Where("id = ?", id).First(&book).Error; err != nil {
		return Book{}, err
	}

	s.DB.Debug().Model(&book).Updates(update)
	return book, nil
}

// DeleteByIDBooks := delete by id
func (s *Server) DeleteByIDBooks(id int) error {
	var book Book
	if err := s.DB.Debug().Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	s.DB.Delete(&book)
	return nil
} 