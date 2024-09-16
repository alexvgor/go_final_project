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
	var env_file string
	if flag.Lookup("test.v") == nil {
		env_file = ".env"
	} else {
		env_file = "../.env"
	}
	err := godotenv.Load(env_file)
	if err != nil {
		if _, ignore_env_file := os.LookupEnv("IGNORE_ENV_FILE"); !ignore_env_file {
			log.Fatal("Error loading .env file (use IGNORE_ENV_FILE variable to ignore this error)")
		} else {
			log.Println("Error loading .env file (was ignored)")
		}
	}
}
