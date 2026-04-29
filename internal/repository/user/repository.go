package userrepo

import (
	"context"
	"database/sql"

	"steam-tracker/internal/model"
)

type Repository struct {
	db *sql.DB
}

type UserRepository interface {
	GetByAPIKey(ctx context.Context, apiKey string) (*model.User, error)
}

var _ UserRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
