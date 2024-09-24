package handlers

import (
	"errors"
	"net/http"
)

type UnroutedHandler struct {
}

func NewUnroutedHandler() *UnroutedHandler {
	return &UnroutedHandler{}
}

func (h *UnroutedHandler) NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondError(w, errors.New("не найдено"), http.StatusNotFound)
	}
}

func (h *UnroutedHandler) MethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondError(w, errors.New("метод не поддерживается"), http.StatusMethodNotAllowed)
	}
}
