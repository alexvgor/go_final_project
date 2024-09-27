package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PublicRoutes(router chi.Router) {
	router.Handle("/*", http.FileServer(http.Dir("./web")))
	router.Route("/api/nextdate", NextDate)
	router.Route("/api/signin", SignIn)
}
