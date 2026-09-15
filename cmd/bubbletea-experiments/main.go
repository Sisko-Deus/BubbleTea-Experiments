package main

import (
	"bubbletea-experiments/internal/logger"
	"bubbletea-experiments/internal/tui"
	"log/slog"

	tea "charm.land/bubbletea/v2"
)

func main() {
	logger.InitLogger()
	slog.Info("Start")

	app := tea.NewProgram(tui.NewRootWindow())
	if _, err := app.Run(); err != nil {
		slog.Error("BubbleTea Main Process. Start ERROR.", slog.String("error", err.Error()))
	}
}
