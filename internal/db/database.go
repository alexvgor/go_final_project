package db

import (
	"fmt"
	"log/slog"
	"os"

	"database/sql"

	"github.com/alexvgor/go_final_project/internal/setup"
	_ "modernc.org/sqlite"
)

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

func CreateDbConnection() (*sql.DB, error) {

	dbFile := setup.GetDbPath()
	slog.Debug(fmt.Sprintf("db file path - %s", dbFile))

	_, err := os.Stat(dbFile)
	var shouldInitDB bool
	if err != nil {
		shouldInitDB = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to open db file - %s", err.Error()))
		return nil, err
	}

	if shouldInitDB {
		err = create_table(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func CloseDbConnection(db *sql.DB) error {
	return db.Close()
}
