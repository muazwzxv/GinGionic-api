package main

import (
	"Go-Learn-API/controllers"
	"Go-Learn-API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.Create)

	router.Run()
}
