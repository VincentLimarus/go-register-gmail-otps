package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}