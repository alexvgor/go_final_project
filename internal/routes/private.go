package routes

import (
	session "github.com/alexvgor/go_final_project/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func PrivateRoutes(router chi.Router) {
	private_router := router.With(session.Session.Validate)
	private_router.Route("/api/task", Task)
	private_router.Route("/api/tasks", Tasks)
}
