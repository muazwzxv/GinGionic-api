package main

import (
	"Go-Learn-API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.Run()
}
