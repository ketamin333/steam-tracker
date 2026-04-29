package trackedgameservice

import (
	"context"

	"steam-tracker/internal/model"
)

func (s *Service) Update(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	return s.repo.Update(ctx, user, gameID, price)
}
