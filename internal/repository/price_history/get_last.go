package pricehistoryrepo

import (
	"context"
	"steam-tracker/internal/model"
)

func (r *Repository) GetLast(ctx context.Context, gameID int, lang string) (model.PriceHistory, error) {
	var ph model.PriceHistory

	err := r.db.QueryRowContext(ctx,
		`SELECT id, game_id, lang, price, currency, discount_percent, checked_at 
		FROM price_history WHERE game_id = $1 AND lang = $2
		ORDER BY checked_at DESC LIMIT 1`,
		gameID, lang,
	).Scan(&ph.ID, &ph.GameID, &ph.Lang, &ph.Price, &ph.Currency, &ph.DiscountPercent, &ph.CheckedAt)

	return ph, err
}
