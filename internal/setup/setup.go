package setup

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	LoadEnv()
	SetLogLevel(os.Getenv("LOG_LEVEL"))
	slog.Info("setup was completed")
}

func GetPort() int {
	port := os.Getenv("TODO_PORT")
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		portNumber = Port
		slog.Warn(fmt.Sprintf("invalid port number was provided - %s (will be used default one)", port))
	}
	return portNumber
}

func GetDbPath() string {
	dbFile := os.Getenv("TODO_DBFILE")
	if len(dbFile) == 0 {
		appPath, err := os.Executable()
		if err != nil {
			slog.Error(fmt.Sprintf("Unable to get executable path - %s", err.Error()))
			os.Exit(1)
		}
		dbFile = filepath.Join(filepath.Dir(appPath), "scheduler.db")
	}
	return dbFile
}

func GetSessionSecret() string {
	secret := os.Getenv("TODO_SECRET")
	if len(secret) == 0 {
		secret = Secret
	}
	return secret
}

func GetSessionPassword() string {
	return os.Getenv("TODO_PASSWORD")
}
