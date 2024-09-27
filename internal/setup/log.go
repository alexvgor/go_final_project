package setup

import (
	"log"
	"log/slog"
	"strings"
)

var customLevels = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func SetLogLevel(level string) {
	if len(level) == 0 {
		level = LogLevel
	}
	slectedLevel, levelExists := customLevels[strings.ToUpper(level)]
	if !levelExists {
		log.Fatal("invalid logging level was provided - ", level)
	}
	slog.SetLogLoggerLevel(slectedLevel)
}
