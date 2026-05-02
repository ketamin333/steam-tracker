package userservice

import (
	"context"
	"steam-tracker/internal/model"
)

func (s *Service) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.Create(ctx, user)
}
