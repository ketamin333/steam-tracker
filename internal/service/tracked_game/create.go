package trackedgameservice

import (
	"context"

	"steam-tracker/internal/model"
)

func (s *Service) Create(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	return s.repo.Create(ctx, user, gameID, price)
}
