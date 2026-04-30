package pricehistoryrepo

import (
	"context"
	"database/sql"

	"steam-tracker/internal/model"
)

type Repository struct {
	db *sql.DB
}

type PriceHistoryRepository interface {
	GetLast(ctx context.Context, gameID int, lang string) (model.PriceHistory, error)
	Create(ctx context.Context, ph *model.PriceHistory) (*model.PriceHistory, error)
	GetLastForGames(ctx context.Context, gameIDs []int, lang string) ([]model.PriceHistory, error)
}

var _ PriceHistoryRepository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}
