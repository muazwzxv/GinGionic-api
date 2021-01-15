package main

import (
	"Go-example/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.Run()
}
