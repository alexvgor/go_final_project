package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alexvgor/go_final_project/internal/routes"
	"github.com/alexvgor/go_final_project/internal/setup"
	"github.com/alexvgor/go_final_project/internal/taskmanager"
)

func main() {

	// close db connection defer
	defer taskmanager.TaskManager.Close()

	router := chi.NewRouter()
	routes.PublicRoutes(router)
	routes.PrivateRoutes(router)
	routes.Unrouted(router)

	port := setup.GetPort()

	slog.Info(fmt.Sprintf("starting app on %d port", port))
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router); err != nil {
		slog.Error(fmt.Sprintf("app was down due to error - %s", err.Error()))
	}
}
