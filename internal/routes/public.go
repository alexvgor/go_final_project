package routes

import (
	"net/http"

	"github.com/alexvgor/go_final_project/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func PublicRoutes(router *chi.Mux) {
	router.Handle("/", http.FileServer(http.Dir("./web")))

	nextDateHandler := handlers.NewNextDateHandler()
	router.HandleFunc("/api/nextdate", nextDateHandler.Handle())
}
