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

	dbConnection, err := db.CreateDbConnection()
	if err != nil {
		slog.Error(fmt.Sprintf("db connection was not created due to error - %s", err.Error()))
		os.Exit(1)
	} else {
		slog.Info("db connection was created")
		defer dbConnection.Close()
	}

	router := chi.NewRouter()
	routes.PublicRoutes(router)

	slog.Info(fmt.Sprintf("starting app on %d port", port))
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router); err != nil {
		slog.Error(fmt.Sprintf("app was down due to error - %s", err.Error()))
		os.Exit(1)
	}
}
