package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/alexvgor/go_final_project/internal/models"
	"github.com/alexvgor/go_final_project/internal/taskmanager"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var taskDTO models.Task
		err := json.NewDecoder(r.Body).Decode(&taskDTO)
		if err != nil {
			RespondErrorUnableToCreateTask(w, errors.New("ошибка десериализации JSON задачи"))
			return
		}

		task_id, err := taskmanager.TaskManager.CreateTask(&taskDTO)
		if err != nil {
			RespondErrorUnableToCreateTask(w, err)
			return
		}

		Respond(w, models.Response{Id: task_id})
	}
}

func (h *TaskHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id_string := r.URL.Query().Get("id")
		if len(id_string) == 0 {
			RespondErrorUnableToGetTask(w, errors.New("не указан идентификатор задачи"))
			return
		}
		id, err := strconv.ParseInt(id_string, 10, 64)
		if err != nil {
			RespondErrorUnableToGetTask(w, errors.New("идентификатор задачи указан в неверном формате"))
			return
		}

		task, err := taskmanager.TaskManager.GetTask(id)
		if err != nil {
			RespondErrorUnableToGetTask(w, err)
			return
		}

		Respond(w, task)
	}
}
