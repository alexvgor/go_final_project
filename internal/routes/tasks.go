package routes

import (
	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Tasks(router chi.Router) {
	tasksHandler := handlers.NewTasksHandler()
	router.Get("/", tasksHandler.Get())
}