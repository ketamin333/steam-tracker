package pricehistoryrepo

import (
	"context"

	"go.rest.api/internal/model"
)

func (r *Repository) Create(ctx context.Context, ph *model.PriceHistory) (*model.PriceHistory, error) {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO price_history (game_id, lang, price, currency, discount_percent) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, game_id, lang, price, currency, discount_percent, checked_at`,
		ph.GameID, ph.Lang, ph.Price, ph.Currency, ph.DiscountPercent,
	).Scan(&ph.ID, &ph.GameID, &ph.Lang, &ph.Price, &ph.Currency, &ph.DiscountPercent, &ph.CheckedAt)

	if err != nil {
		return nil, err
	}

	return ph, nil
}
