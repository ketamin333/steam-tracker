package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"go.rest.api/internal/app"
)

func main() {
	_ = godotenv.Load()

	srv, err := app.Bootstrap()

	if err != nil {
		slog.Error("bootstrap failed", "err", err)
		os.Exit(1)
	}

	if err := srv.Run(); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
