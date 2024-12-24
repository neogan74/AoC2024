package logger

import (
	"log/slog"
	"os"
)

const (
	_envDebug = "debug"
	_envDev   = "dev"
	_envProd  = "prod"
)

type DefaultLoggerInterface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type Logger interface {
	DefaultLoggerInterface
}

func NewLogger(env string) Logger {
	var log slog.Logger

	switch env {
	case _envDebug:
		log = *slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case _envDev:
		log = *slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case _envProd:
		log = *slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
