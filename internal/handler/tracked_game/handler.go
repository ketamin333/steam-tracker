package trackedgamehandler

import "context"

type Handler struct {
	svc TrackedGameService
}

type TrackedGameService interface {
	Track(ctx context.Context, userID, gameID int) error
	Untrack(ctx context.Context, userID, gameID int) error
}

func New(svc TrackedGameService) *Handler {
	return &Handler{svc: svc}
}
