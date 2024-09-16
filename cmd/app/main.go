package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alexvgor/go_final_project/pkg/setup"
	"github.com/alexvgor/go_final_project/tests"
)

func main() {

	setup.LoadEnv()

	port := os.Getenv("TODO_PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		portNumber = tests.Port
		log.Printf("invalid port number was provided - %s (will be used default one)\n", port)
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Printf("starting app on %d port ... \n", portNumber)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", portNumber), nil)
	if err != nil {
		log.Fatal(err)
	}
}
