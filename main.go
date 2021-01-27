package main

import (
	"Go-Learn-API/controllers"
	"Go-Learn-API/middleware"
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
	router.POST("/books", middleware.TokenAuthMiddleware(), controllers.CreateBooks)
	router.GET("/books", middleware.TokenAuthMiddleware(), controllers.GetAllBooks)
	router.GET("/books/:id", middleware.TokenAuthMiddleware(), controllers.GetByIDBooks)
	router.PATCH("/books/:id", middleware.TokenAuthMiddleware(), controllers.UpdateByIDBooks)
	router.DELETE("books/:id", middleware.TokenAuthMiddleware(), controllers.DeleteByIDBooks)

	router.Run()
}
