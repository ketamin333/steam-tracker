package userhandler

import (
	"context"
	"steam-tracker/internal/model"
)

type Handler struct {
	svc UserService
}

type UserService interface {
	Update(ctx context.Context, user *model.User, email *string, lang *string) (*model.User, error)
}

func New(svc UserService) *Handler {
	return &Handler{svc: svc}
}
