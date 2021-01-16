package controllers

import (
	"Go-Learn-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBookDTO := DTO new book
type CreateBookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookDTO := DTO update book
type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetAllBooks := returns all
func GetAllBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"Data": books})
}

// CreateBooks := create new book
func CreateBooks(ctx *gin.Context) {
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

// GetByIDBooks := returns by ID
func GetByIDBooks(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateByIDBooks := update by ID
func UpdateByIDBooks(ctx *gin.Context) {
	var book models.Book

	// Query data
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var validate UpdateBookDTO
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(validate)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}
