package main

import (
	"Go-Learn-API/controllers"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func viperEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("invalid type insertion")
	}

	return value
}

func main() {
	router := gin.Default()

	fmt.Println(viperEnv("DBPASSWORD"))

	// models.ConnectDB()

	router.GET("/books", controllers.FindBooks)

	// router.Run()
}
