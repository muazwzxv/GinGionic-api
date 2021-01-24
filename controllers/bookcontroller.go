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
	books, err := models.Model.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erorr": err})
	}

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

	_, err := models.Model.CreateBooks(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// GetByIDBooks := returns by ID
func GetByIDBooks(ctx *gin.Context) {
	if book, err := models.Model.GetByIDBooks(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": book})
	}

	// if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
	// 	return
	// }
}

// UpdateByIDBooks := update by ID
func UpdateByIDBooks(ctx *gin.Context) {
	var book models.Book

	// Query data
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not exist"})
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

// DeleteByIDBooks := delete by id
func DeleteByIDBooks(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record does not exist", "deleted": false})
		return
	}

	models.Models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}
