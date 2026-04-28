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

func (s *Service) Create(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	return s.repo.Create(ctx, user, gameID, price)
}

func (s *Service) Delete(ctx context.Context, user *model.User, gameID int) error {
	return s.repo.Delete(ctx, user, gameID)
}

func (s *Service) Update(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	return s.repo.Update(ctx, user, gameID, price)
}
