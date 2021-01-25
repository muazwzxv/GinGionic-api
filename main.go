package main

import (
	"Go-Learn-API/controllers"
	"Go-Learn-API/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	_, err := models.Model.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the databaase: ", err)
	}

	// Auth endpoints
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	// Books endpoint
	router.POST("/books", controllers.CreateBooks)
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetByIDBooks)
	router.PATCH("/books/:id", controllers.UpdateByIDBooks)
	router.DELETE("books/:id", controllers.DeleteByIDBooks)

	router.Run()
}
