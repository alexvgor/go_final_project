package handlers

import (
	"encoding/json"
	"net/http"

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
			RespondErrorInvalidJson(w)
			return
		}

		task_id, err := taskmanager.TaskManager.CreateTask(&taskDTO)
		if err != nil {
			RespondError(w, err)
			return
		}

		Respond(w, models.Response{Id: task_id})
	}
}
