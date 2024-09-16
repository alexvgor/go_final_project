package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/alexvgor/go_final_project/internal/db"
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

	dbConnection := db.CreateDbConnection()
	if dbConnection != nil {
		slog.Info("db connection was created")
		defer dbConnection.Close()
	} else {
		slog.Error("db connection was not created")
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir("./web")))
	slog.Info(fmt.Sprintf("starting app on %d port", portNumber))
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", portNumber), nil)
	if err != nil {
		slog.Error(fmt.Sprintf("App was down due to error - %s", err.Error()))
		os.Exit(1)
	}
}
