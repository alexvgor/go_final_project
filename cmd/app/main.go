package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/alexvgor/go_final_project/pkg/setup"
	"github.com/alexvgor/go_final_project/tests"
)

func main() {

	setup.LoadEnv()

	setup.SetLogLevel(os.Getenv("LOG_LEVEL"))

	port := os.Getenv("TODO_PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		portNumber = tests.Port
		slog.Warn(fmt.Sprintf("invalid port number was provided - %s (will be used default one)", port))
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))
	slog.Info(fmt.Sprintf("starting app on %d port", portNumber))
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", portNumber), nil)
	if err != nil {
		slog.Error(fmt.Sprintf("App was down due to error - %s", err.Error()))
		log.Fatal(err)
	}
}
