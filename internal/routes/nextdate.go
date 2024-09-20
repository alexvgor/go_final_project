package routes

import (
	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NextDate(router chi.Router) {
	nextDateHandler := handlers.NewNextDateHandler()
	router.Get("/", nextDateHandler.Get())
}
