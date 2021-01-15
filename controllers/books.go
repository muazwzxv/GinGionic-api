package controllers

import (
	"Go-Learn-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"Data": books})
}
