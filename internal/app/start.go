package app

import (
	"context"
	"log/slog"
	"time"
)

func (a *App) Start() error {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := a.Tracker.Run(context.Background()); err != nil {
				slog.Error("app tracker run failed", "err", err)
			}
		}
	}()

	return a.Server.Run()
}
