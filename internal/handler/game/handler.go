package gamehandler

import (
	"context"

	"go.rest.api/internal/model"
)

type Handler struct {
	svc GameService
}

type GameService interface {
	Search(ctx context.Context, user *model.User, query string) ([]model.Game, error)
}

func New(svc GameService) *Handler {
	return &Handler{svc: svc}
}
