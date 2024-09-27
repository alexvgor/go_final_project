package routes

import (
	session "github.com/alexvgor/go_final_project/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func PrivateRoutes(router chi.Router) {
	privateRouter := router.With(session.Session.Validate)
	privateRouter.Route("/api/task", Task)
	privateRouter.Route("/api/tasks", Tasks)
}
