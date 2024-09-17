package setup

import (
	"log"
	"log/slog"
	"strings"
)

var CUSTOM_LEVELS = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func SetLogLevel(level string) {
	if len(level) == 0 {
		level = "INFO"
	}
	slected_level, level_exists := CUSTOM_LEVELS[strings.ToUpper(level)]
	if !level_exists {
		log.Fatal("Invalid logging level was provided - ", level)
	}
	slog.SetLogLoggerLevel(slected_level)
}
