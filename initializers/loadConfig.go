package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error load .env file")
	}
}
