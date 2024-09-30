package database

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"database/sql"

	"github.com/alexvgor/go_final_project/internal/models"
	"github.com/alexvgor/go_final_project/internal/setup"
	_ "modernc.org/sqlite"
)

type Db struct {
	db *sql.DB
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date CHAR(8) NOT NULL DEFAULT "", title VARCHAR(256) NOT NULL DEFAULT "", comment VARCHAR(256) NOT NULL DEFAULT "", repeat VARCHAR(128) NOT NULL DEFAULT "")`)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to create table - %s", err.Error()))
		return err
	}
	slog.Debug("table 'scheduler' was created")
	err = createIndex(db)
	return err
}

func createIndex(db *sql.DB) error {
	_, err := db.Exec("CREATE INDEX date ON scheduler (id)")
	if err != nil {
		slog.Error(fmt.Sprintf("unable to create index - %s", err.Error()))
		return err
	}
	slog.Debug("column 'date' from table 'scheduler' was indexed")
	return nil
}

func Create() (Db, error) {

	var db Db

	dbFile := setup.GetDbPath()
	slog.Debug(fmt.Sprintf("db file path - %s", dbFile))

	_, err := os.Stat(dbFile)
	var shouldInitDB bool
	if err != nil {
		shouldInitDB = true
	}

	dbConnection, err := sql.Open("sqlite", dbFile)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to open db file - %s", err.Error()))
		return db, err
	}

	if shouldInitDB {
		err = createTable(dbConnection)
		if err != nil {
			return db, err
		}
	}

	db.db = dbConnection

	return db, nil
}

func (db Db) Close() error {
	return db.db.Close()
}

func (db Db) CreateTask(task *models.Task) (int64, error) {
	res, err := db.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)", task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to insert task - %s", err.Error()))
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		slog.Error(fmt.Sprintf("unable to get inserted task id - %s", err.Error()))
		return 0, err
	}
	return id, nil
}

func (db Db) GetTask(id int64) (models.Task, error) {
	var task models.Task
	row := db.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id)
	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, errors.New("задача не найдена")
		}
		slog.Error(fmt.Sprintf("unable to get task by id - %s", err.Error()))
	}
	return task, err
}

func (db Db) DeleteTask(id int64) error {
	_, err := db.db.Exec("DELETE FROM scheduler WHERE id = ?", id)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to delete task by id - %s", err.Error()))
		return errors.New("задача не удалена")
	}
	return nil
}

func (db Db) UpdateTask(task *models.Task) error {
	res, err := db.db.Exec("UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?", task.Date, task.Title, task.Comment, task.Repeat, task.Id)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to update task by id - %s", err.Error()))
		return errors.New("задача не обнавлена")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		slog.Error(fmt.Sprintf("unable to confirm that task was updated - %s", err.Error()))
		return errors.New("ошибка подтверждения изменения")
	} else if rowsAffected == 0 {
		slog.Error("unable to confirm that task was updated")
		return errors.New("обновление задачи не привело к изменениям")
	}
	return nil
}

func (db Db) GetTasks() ([]models.Task, error) {
	row, err := db.db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?", setup.DbQueryLimit)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to get tasks - %s", err.Error()))
		return nil, err
	}
	return parseTaskRows(row)
}

func (db Db) GetTasksFilteredByDate(date string) ([]models.Task, error) {
	row, err := db.db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date LIMIT ?", date, setup.DbQueryLimit)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to get tasks - %s", err.Error()))
		return nil, err
	}
	return parseTaskRows(row)
}

func (db Db) GetTasksFilteredByTitleOrComment(search string) ([]models.Task, error) {
	row, err := db.db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?", search, search, setup.DbQueryLimit)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to get tasks - %s", err.Error()))
		return nil, err
	}
	return parseTaskRows(row)
}

func parseTaskRows(row *sql.Rows) ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	for row.Next() {
		var task models.Task
		err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			slog.Error(fmt.Sprintf("unable to parse tasks row - %s", err.Error()))
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
