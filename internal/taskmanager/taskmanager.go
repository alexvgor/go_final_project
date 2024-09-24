package taskmanager

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/alexvgor/go_final_project/internal/database"
	"github.com/alexvgor/go_final_project/internal/models"
	"github.com/alexvgor/go_final_project/internal/setup"
)

type TaskManagerInstance struct {
	db database.Db
}

var TaskManager TaskManagerInstance

func Init(db database.Db) {
	TaskManager = New(db)
}

func New(db database.Db) TaskManagerInstance {
	tm := TaskManagerInstance{db}
	return tm
}

func parseTasksIdAsString(tasks []models.Task) []models.ResponseTask {
	response_tasks := make([]models.ResponseTask, len(tasks))
	for task_index, task := range tasks {
		response_tasks[task_index] = models.ResponseTask{
			Id:      strconv.FormatInt(task.Id, 10),
			Date:    task.Date,
			Title:   task.Title,
			Comment: task.Comment,
			Repeat:  task.Repeat,
		}
	}
	return response_tasks
}

func (tm TaskManagerInstance) GetTasks() ([]models.ResponseTask, error) {
	tasks, err := tm.db.GetTasks()
	if err != nil {
		return []models.ResponseTask{}, errors.New("ошибка поиска задач")
	}
	return parseTasksIdAsString(tasks), err
}

func (tm TaskManagerInstance) GetTasksFilteredByDate(search string) ([]models.ResponseTask, error) {
	tasks, err := tm.db.GetTasksFilteredByDate(search)
	if err != nil {
		return []models.ResponseTask{}, errors.New("ошибка поиска задач по времени")
	}
	return parseTasksIdAsString(tasks), err
}

func (tm TaskManagerInstance) GetTasksFilteredByTitleOrComment(search string) ([]models.ResponseTask, error) {
	tasks, err := tm.db.GetTasksFilteredByTitleOrComment(search)
	if err != nil {
		return []models.ResponseTask{}, errors.New("ошибка поиска задач по заголовку или комментарию")
	}
	return parseTasksIdAsString(tasks), err
}

func (tm TaskManagerInstance) CreateTask(task *models.Task) (int64, error) {
	validated_task, err := validateTask(task)
	if err != nil {
		return 0, err
	}

	id, err := tm.db.CreateTask(validated_task)
	if err != nil {
		return 0, errors.New("ошибка сохранения новой задачи")
	}

	return id, nil
}

func validateTask(task *models.Task) (*models.Task, error) {
	if task.Title == "" {
		return nil, errors.New("не указан заголовок задачи")
	}

	now := time.Now()
	today := now.Format(setup.ParseDateFormat)
	if len(strings.TrimSpace(task.Date)) == 0 {
		task.Date = today
		return task, nil
	}

	date, err := time.Parse(setup.ParseDateFormat, task.Date)
	if err != nil {
		return nil, errors.New("дата представлена в неверном формате")
	}

	if date.Format(setup.ParseDateFormat) < today {
		if len(strings.TrimSpace(task.Repeat)) == 0 {
			task.Date = today
		} else {
			next_date, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return nil, errors.New("правило повторения указано в неверном формате")
			}
			task.Date = next_date
		}
	}

	return task, nil
}
