package trackedgameservice

import (
	"context"

	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
)

type Service struct {
	repo trackedgamerepo.TrackingRepository
}

func New(repo trackedgamerepo.TrackingRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Track(ctx context.Context, userID, gameID int) error {
	return s.repo.Add(ctx, userID, gameID)
}

func (s *Service) Untrack(ctx context.Context, userID, gameID int) error {
	return s.repo.Remove(ctx, userID, gameID)
}
