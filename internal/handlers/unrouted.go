package handlers

import (
	"net/http"
)

type UnroutedHandler struct {
}

func NewUnroutedHandler() *UnroutedHandler {
	return &UnroutedHandler{}
}

func (h *UnroutedHandler) NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Не найдено", http.StatusNotFound)
	}
}

func (h *UnroutedHandler) MethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
