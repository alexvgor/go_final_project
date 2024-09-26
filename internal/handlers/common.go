package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alexvgor/go_final_project/internal/models"
)

type ErrorsStruct struct {
	UnableToFindNextDate error
	UnableToGetTasks     error
	UnableToCreateTasks  error
	UnableToGetTask      error
	UnableToUpdateTask   error
}

var HandlerErrors = ErrorsStruct{
	UnableToFindNextDate: errors.New("ошибка получения следующей даты задачи"),
	UnableToGetTasks:     errors.New("ошибка получения задач"),
	UnableToCreateTasks:  errors.New("ошибка создания задачи"),
	UnableToGetTask:      errors.New("ошибка получения задачи"),
	UnableToUpdateTask:   errors.New("ошибка изменения задачи"),
}

func Respond(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(v)
}

func RespondJsonError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(models.Response{Error: err.Error()})
}

func RespondErrorUnableToGetTasks(w http.ResponseWriter, err error) {
	RespondJsonError(w, errors.Join(HandlerErrors.UnableToGetTasks, err))
}

func RespondErrorUnableToUpdateTask(w http.ResponseWriter, err error) {
	RespondJsonError(w, errors.Join(HandlerErrors.UnableToUpdateTask, err))
}

func RespondErrorUnableToGetTask(w http.ResponseWriter, err error) {
	RespondJsonError(w, errors.Join(HandlerErrors.UnableToGetTask, err))
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
