package gameservice

import (
	"steam-tracker/internal/client/steam"
	gamerepo "steam-tracker/internal/repository/game"
)

type Service struct {
	steam steam.SteamClient
	repo  gamerepo.GameRepository
}

func New(steam steam.SteamClient, repo gamerepo.GameRepository) *Service {
	return &Service{steam: steam, repo: repo}
}
