package db

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"database/sql"

	_ "modernc.org/sqlite"
)

func CreateTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE scheduler (id INTEGER PRIMARY KEY AUTOINCREMENT, date CHAR(8) NOT NULL DEFAULT "", title VARCHAR(256) NOT NULL DEFAULT "", comment VARCHAR(256) NOT NULL DEFAULT "", repeat VARCHAR(128) NOT NULL DEFAULT "")`)
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to create table - %s", err.Error()))
		os.Exit(1)
	}
	slog.Debug("table 'scheduler' was created")
}

func CreateIndex(db *sql.DB) {
	_, err := db.Exec("CREATE INDEX date ON scheduler (id)")
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to create index - %s", err.Error()))
		os.Exit(1)
	}
	slog.Debug("column 'date' from table 'scheduler' was indexed")
}

func CreateDbConnection() *sql.DB {
	appPath, err := os.Executable()
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get executable path - %s", err.Error()))
		os.Exit(1)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	slog.Debug(fmt.Sprintf("db file path - %s", dbFile))

	_, err = os.Stat(dbFile)
	var shouldInitDB bool
	if err != nil {
		shouldInitDB = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to open db file - %s", err.Error()))
		os.Exit(1)
	}

	if shouldInitDB {
		CreateTable(db)
		CreateIndex(db)
	}

	return db
}

func CloseDbConnection(db *sql.DB) {
	db.Close()
}
