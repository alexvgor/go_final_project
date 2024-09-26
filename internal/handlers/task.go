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

func getQueryId(r *http.Request) (int64, error) {
	id_string := r.URL.Query().Get("id")
	if len(id_string) == 0 {
		return 0, errors.New("не указан идентификатор задачи")
	}
	id, err := strconv.ParseInt(id_string, 10, 64)
	if err != nil {
		return id, errors.New("идентификатор задачи указан в неверном формате")
	}
	return id, nil
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
		id, err := getQueryId(r)
		if err != nil {
			RespondErrorUnableToGetTask(w, err)
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

func (h *TaskHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var taskDTO models.ResponseTask
		err := json.NewDecoder(r.Body).Decode(&taskDTO)
		if err != nil {
			RespondErrorUnableToUpdateTask(w, errors.New("ошибка десериализации JSON задачи"))
			return
		}
		id, err := strconv.ParseInt(taskDTO.Id, 10, 64)
		if err != nil {
			RespondErrorUnableToUpdateTask(w, errors.New("идентификатор задачи указан в неверном формате"))
			return
		}
		task := models.Task{
			Id:      id,
			Date:    taskDTO.Date,
			Title:   taskDTO.Title,
			Comment: taskDTO.Comment,
			Repeat:  taskDTO.Repeat,
		}
		err = taskmanager.TaskManager.UpdateTask(&task)
		if err != nil {
			RespondErrorUnableToUpdateTask(w, err)
			return
		}
		Respond(w, models.Response{})
	}
}

func (h *TaskHandler) PostDone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getQueryId(r)
		if err != nil {
			RespondErrorUnableToSetTaskAsDone(w, err)
			return
		}

		err = taskmanager.TaskManager.SetTaskAsDone(id)
		if err != nil {
			RespondErrorUnableToSetTaskAsDone(w, err)
			return
		}

		Respond(w, models.Response{})
	}
}
