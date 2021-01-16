package controllers

import (
	"Go-Learn-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBookDTO := DTO for new book
type CreateBookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// FindBooks returns all books
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"Data": books})
}

// Create := create new book
func Create(ctx *gin.Context) {
	var input CreateBookDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}
