package app

import (
	"go.rest.api/internal/db"
	game_handler "go.rest.api/internal/handlers/game"
	game_repository "go.rest.api/internal/repository/game"
	"go.rest.api/internal/router"
	"go.rest.api/internal/server"
)

func Bootstrap() *server.Server {
	repo := game_repository.NewGameRepository(db.DB)
	gameHandler := game_handler.NewHandler(repo)

	r := router.Setup(gameHandler)
	return server.NewServer(":8080", r)
}
