package routes

import (
	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Task(router chi.Router) {
	taskHandler := handlers.NewTaskHandler()
	router.Post("/", taskHandler.Post())
	router.Get("/", taskHandler.Get())
	router.Put("/", taskHandler.Put())
}
