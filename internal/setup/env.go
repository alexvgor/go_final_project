package setup

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var envFile string
	if testing.Testing() {
		envFile = "../.env"
	} else {
		envFile = ".env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Error loading .env file")
	}
}
