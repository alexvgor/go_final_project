package setup

/*
	package is used to load and parse .env file
*/

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	var envFile string
	if flag.Lookup("test.v") == nil {
		envFile = ".env"
	} else {
		envFile = "../.env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		if _, ignoreEnvFile := os.LookupEnv("IGNORE_ENV_FILE"); !ignoreEnvFile {
			log.Fatal("Error loading .env file (use IGNORE_ENV_FILE variable to ignore this error)")
		} else {
			log.Println("Error loading .env file (was ignored)")
		}
	}
}
