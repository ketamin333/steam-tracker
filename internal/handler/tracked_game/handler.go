package trackedgamehandler

import (
	"context"

	"steam-tracker/internal/model"
)

type Handler struct {
	svc TrackedGameService
}

type TrackedGameService interface {
	Create(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error)
	Update(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error)
	Delete(ctx context.Context, user *model.User, gameID int) error
}

func New(svc TrackedGameService) *Handler {
	return &Handler{svc: svc}
}
