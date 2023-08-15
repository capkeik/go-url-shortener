package main

import (
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	localEnv = "local"
	prodEnv  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	logger.Info("starting url-shortener")
	logger.Debug("debug messages are enabled")
	fmt.Println(cfg)
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case localEnv:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case prodEnv:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
