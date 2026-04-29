package gamerepo

import (
	"context"
	"database/sql"

	"steam-tracker/internal/model"
)

type Repository struct {
	db *sql.DB
}

type GameRepository interface {
	Upsert(ctx context.Context, g *model.Game) error
}

var _ GameRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
