package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env doesnt exist")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}