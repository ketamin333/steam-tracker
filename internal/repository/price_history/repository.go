package pricehistoryrepo

import (
	"context"
	"database/sql"

	"go.rest.api/internal/model"
)

type Repository struct {
	db *sql.DB
}

type PriceHistoryRepository interface {
	Create(ctx context.Context, ph *model.PriceHistory) (*model.PriceHistory, error)
}

var _ PriceHistoryRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
