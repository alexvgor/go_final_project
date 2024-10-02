package routes

import (
	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Unrouted(router chi.Router) {
	unroutedHandler := handlers.NewUnroutedHandler()
	router.NotFound(unroutedHandler.NotFound())
	router.MethodNotAllowed(unroutedHandler.MethodNotAllowed())
}
