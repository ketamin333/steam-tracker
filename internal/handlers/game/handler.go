package game_handler

import (
	"go.rest.api/internal/client"
	gamerepository "go.rest.api/internal/repository/game"
)

type Handler struct {
	repo   *gamerepository.GameRepository
	client *client.SteamClient
}

func NewGameHandler(repo *gamerepository.GameRepository, client *client.SteamClient) *Handler {
	return &Handler{repo: repo, client: client}
}
