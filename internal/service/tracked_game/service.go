package trackedgameservice

import (
	"context"

	"go.rest.api/internal/model"
	trackedgamerepo "go.rest.api/internal/repository/tracked_game"
)

type Service struct {
	repo trackedgamerepo.TrackingRepository
}

func New(repo trackedgamerepo.TrackingRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Track(ctx context.Context, userID, gameID int, price *float64) (*model.TrackedGame, error) {
	return s.repo.Add(ctx, userID, gameID, price)
}

func (s *Service) Untrack(ctx context.Context, userID, gameID int) error {
	return s.repo.Remove(ctx, userID, gameID)
}
