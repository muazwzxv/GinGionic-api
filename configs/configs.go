package configs

import (
	"errors"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

// DatabaseConfig := database config
type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

// Reads Env
func readEnv(key string) string {
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

// GetTokenSecret := returns jwt secret
func GetTokenSecret() string {
	return readEnv("API_SECRET")
}

// DBConfig returns database config
func DBConfig() (DatabaseConfig, error) {
	port, err := strconv.Atoi(readEnv("DBPORT"))
	if err != nil {
		return DatabaseConfig{}, errors.New("Error converting string to int")
	}

	config := &DatabaseConfig{
		User:     readEnv("DBUSER"),
		Password: readEnv("DBPASSWORD"),
		Host:     readEnv("DBHOST"),
		DBName:   readEnv("DBNAME"),
		Port:     port,
	}

	return *config, nil
}
