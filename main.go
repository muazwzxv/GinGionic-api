package main

import (
	"Go-Learn-API/controllers"
	"Go-Learn-API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.POST("/books", controllers.CreateBooks)
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetByIDBooks)
	router.PATCH("/books/:id", controllers.UpdateByIDBooks)
	router.DELETE("books/:id", controllers.DeleteByIDBooks)

	router.Run()
}
