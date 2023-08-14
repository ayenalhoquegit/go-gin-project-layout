package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		log.Println("Error loading .env file")
		log.Fatal(err)
	}
	dbUser := GetEnvValue("DB_USER")
	log.Println(dbUser)
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
