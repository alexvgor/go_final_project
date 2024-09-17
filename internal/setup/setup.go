package setup

import (
	"log/slog"
	"os"
)

func Init() {
	LoadEnv()
	SetLogLevel(os.Getenv("LOG_LEVEL"))
	slog.Info("setup was completed")
}
