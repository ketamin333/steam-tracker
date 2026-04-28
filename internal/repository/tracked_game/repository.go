package trackedgamerepo

import (
	"context"
	"database/sql"

	"go.rest.api/internal/model"
)

type Repository struct {
	db *sql.DB
}

type TrackingRepository interface {
	Add(ctx context.Context, userID, gameID int, price *float64) (*model.TrackedGame, error)
	Remove(ctx context.Context, userID, gameID int) error
}

var _ TrackingRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
