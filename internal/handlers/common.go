package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexvgor/go_final_project/internal/models"
)

type ErrorsStruct struct {
	InvalidJson error
}

var HandlerErrors = ErrorsStruct{
	InvalidJson: errors.New("ошибка десериализации JSON"),
}

func Respond(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(v)
}

func RespondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(models.Response{Error: err.Error()})
}

func RespondErrorInvalidJson(w http.ResponseWriter) {
	RespondError(w, HandlerErrors.InvalidJson)
}
