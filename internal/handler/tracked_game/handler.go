package trackedgamehandler

import (
	"context"

	"go.rest.api/internal/model"
)

type Handler struct {
	svc TrackedGameService
}

type TrackedGameService interface {
	Track(ctx context.Context, userID, gameID int, price *float64) (*model.TrackedGame, error)
	Untrack(ctx context.Context, userID, gameID int) error
}

func New(svc TrackedGameService) *Handler {
	return &Handler{svc: svc}
}
