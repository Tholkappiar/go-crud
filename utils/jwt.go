package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SECRET_KEY string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatal("No Database URL.")
	}
}
