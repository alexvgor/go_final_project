package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexvgor/go_final_project/internal/models"
)

type ErrorsStruct struct {
	UnableToFindNextDate error
	InvalidJson          error

	UnableToGetTasks    error
	UnableToCreateTasks error
}

var HandlerErrors = ErrorsStruct{
	UnableToFindNextDate: errors.New("ошибка получения следующей даты задачи"),
	InvalidJson:          errors.New("ошибка десериализации JSON"),
	UnableToGetTasks:     errors.New("ошибка получения задач"),
	UnableToCreateTasks:  errors.New("ошибка создания задачи"),
}

func Respond(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(v)
}

func RespondJsonError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(models.Response{Error: err.Error()})
}

func RespondErrorInvalidJson(w http.ResponseWriter) {
	RespondJsonError(w, HandlerErrors.InvalidJson)
}

func RespondErrorUnableToGetTasks(w http.ResponseWriter, err error) {
	RespondJsonError(w, errors.Join(HandlerErrors.UnableToGetTasks, err))
}

func RespondErrorUnableToCreateTask(w http.ResponseWriter, err error) {
	RespondJsonError(w, errors.Join(HandlerErrors.UnableToCreateTasks, err))
}

func RespondError(w http.ResponseWriter, err error, code int) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	http.Error(w, err.Error(), code)
}

func RespondErrorUnableToFindNextDate(w http.ResponseWriter, err error) {
	RespondError(w, errors.Join(HandlerErrors.UnableToFindNextDate, err), http.StatusBadRequest)
}
