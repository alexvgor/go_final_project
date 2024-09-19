package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/alexvgor/go_final_project/internal/db"
	"github.com/alexvgor/go_final_project/internal/routes"
	"github.com/alexvgor/go_final_project/internal/setup"
)

func main() {

	setup.Init()

	port := setup.GetPort()

	dbConnection := db.CreateDbConnection()
	if dbConnection != nil {
		slog.Info("db connection was created")
		defer dbConnection.Close()
	} else {
		slog.Error("db connection was not created")
		os.Exit(1)
	}

	router := chi.NewRouter()
	routes.PublicRoutes(router)

	slog.Info(fmt.Sprintf("starting app on %d port", port))
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router); err != nil {
		slog.Error(fmt.Sprintf("App was down due to error - %s", err.Error()))
		os.Exit(1)
	}
}
