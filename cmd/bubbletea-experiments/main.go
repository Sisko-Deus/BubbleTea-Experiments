package main

import (
	"bubbletea-experiments/internal/logger"
	"log/slog"
)

func main() {
	logger.InitLogger()
	slog.Info("Start")
}
