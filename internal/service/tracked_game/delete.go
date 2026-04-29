package trackedgameservice

import (
	"context"

	"go.rest.api/internal/model"
)

func (s *Service) Delete(ctx context.Context, user *model.User, gameID int) error {
	return s.repo.Delete(ctx, user, gameID)
}
