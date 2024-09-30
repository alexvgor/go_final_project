package routes

import (
	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SignIn(router chi.Router) {
	authHandler := handlers.NewAuthHandler()
	router.Post("/", authHandler.SignIn())
}
