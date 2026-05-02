package userservice

import (
	"context"
	"steam-tracker/internal/model"
)

func (s *Service) Update(ctx context.Context, user *model.User, email *string, lang *string) (*model.User, error) {
	return s.repo.Update(ctx, user, email, lang)
}
