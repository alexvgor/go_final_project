package database

import (
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

func create_table(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date CHAR(8) NOT NULL DEFAULT "", title VARCHAR(256) NOT NULL DEFAULT "", comment VARCHAR(256) NOT NULL DEFAULT "", repeat VARCHAR(128) NOT NULL DEFAULT "")`)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to create table - %s", err.Error()))
		return err
	}
	slog.Debug("table 'scheduler' was created")
	err = create_index(db)
	return err
}

func create_index(db *sql.DB) error {
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

	db_connection, err := sql.Open("sqlite", dbFile)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to open db file - %s", err.Error()))
		return db, err
	}

	if shouldInitDB {
		err = create_table(db_connection)
		if err != nil {
			return db, err
		}
	}

	db.db = db_connection

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
