package trackedgameservice

import (
	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
)

type Service struct {
	repo trackedgamerepo.TrackingRepository
}

func New(repo trackedgamerepo.TrackingRepository) *Service {
	return &Service{repo: repo}
}
