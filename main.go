package main

import (
	"log/slog"
	"os"

	"steam-tracker/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	srv, err := app.Bootstrap()

	if err != nil {
		slog.Error("bootstrap failed", "err", err)
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
