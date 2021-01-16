package models

import (
	"Go-Learn-API/configs"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// mysql dialect in gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// DB instance
var DB *gorm.DB

func viperEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("invalid type insertion")
	}

	return value
}

// ConnectDB connects to the database
func ConnectDB() {

	dbConfig, err := configs.DBConfig()

	if err != nil {
		panic("Failed to fetch database")
	}

	// db, err := gorm.Open("mysql", "user:password@(localhost:3306)/gormLearn?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
		),
	)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})
	defer db.Close()

	DB = db
}
