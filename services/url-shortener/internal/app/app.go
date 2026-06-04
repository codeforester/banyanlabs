package app

import (
	"context"
	"log/slog"
	"time"
)

const serviceName = "url-shortener"

type App struct {
	logger  *slog.Logger
	started time.Time
}

type Options struct {
	Logger *slog.Logger
}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
}

func New(options Options) *App {
	logger := options.Logger
	if logger == nil {
		logger = slog.Default()
	}

	return &App{
		logger:  logger.With(slog.String("component", "app")),
		started: time.Now().UTC(),
	}
}

func (app *App) Health(_ context.Context) Health {
	return Health{
		Service: serviceName,
		Status:  "ok",
		Uptime:  time.Since(app.started).Round(time.Millisecond).String(),
	}
}
