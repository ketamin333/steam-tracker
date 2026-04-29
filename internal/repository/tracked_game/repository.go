package trackedgamerepo

import (
	"context"
	"database/sql"

	"steam-tracker/internal/model"
)

type Repository struct {
	db *sql.DB
}

type TrackingRepository interface {
	GetAll(ctx context.Context) ([]TrackedGameRow, error)
	Create(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error)
	Update(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error)
	Delete(ctx context.Context, user *model.User, gameID int) error
}

var _ TrackingRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
