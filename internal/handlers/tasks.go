package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexvgor/go_final_project/internal/models"
	"github.com/alexvgor/go_final_project/internal/setup"
	"github.com/alexvgor/go_final_project/internal/taskmanager"
)

type TasksHandler struct {
}

func NewTasksHandler() *TasksHandler {
	return &TasksHandler{}
}

func (h *TasksHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		var tasks []models.ResponseTask
		var err error
		if len(search) == 0 {
			tasks, err = taskmanager.TaskManager.GetTasks()
		} else {
			date, dateErr := time.Parse("02.01.2006", search)
			if dateErr == nil {
				tasks, err = taskmanager.TaskManager.GetTasksFilteredByDate(date.Format(setup.ParseDateFormat))
			} else {
				tasks, err = taskmanager.TaskManager.GetTasksFilteredByTitleOrComment(fmt.Sprintf("%%%s%%", search))
			}
		}
		if err != nil {
			RespondErrorUnableToGetTasks(w, err)
			return
		}

		Respond(w, models.ResponseTasks{Tasks: tasks})
	}
}
