package gameservice

import (
	"go.rest.api/internal/client/steam"
	gamerepo "go.rest.api/internal/repository/game"
)

type Service struct {
	steam steam.SteamClient
	repo  gamerepo.GameRepository
}

func New(steam steam.SteamClient, repo gamerepo.GameRepository) *Service {
	return &Service{
		steam: steam,
		repo:  repo,
	}
}
