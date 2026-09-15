package logger

import (
	"log/slog"
	"os"
)

func InitLogger() {
	writer := os.Stdout
	logger := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	slog.Info("Logger. Initialized.")
}
