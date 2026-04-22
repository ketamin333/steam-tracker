package game_handler

import "go.rest.api/internal/repository/game"

type Handler struct {
	repo *game_repository.GameRepository
}

func NewHandler(repo *game_repository.GameRepository) *Handler {
	return &Handler{repo: repo}
}
