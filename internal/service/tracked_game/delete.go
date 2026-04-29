package trackedgameservice

import (
	"context"

	"steam-tracker/internal/model"
)

func (s *Service) Delete(ctx context.Context, user *model.User, gameID int) error {
	return s.repo.Delete(ctx, user, gameID)
}
