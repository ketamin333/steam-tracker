package pricehistoryrepo

import (
	"context"
	"steam-tracker/internal/model"

	"github.com/lib/pq"
)

func (r *Repository) GetLastForGames(ctx context.Context, gameIDs []int, lang string) ([]model.PriceHistory, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT DISTINCT ON (game_id) id, game_id, lang, price, currency, discount_percent, checked_at
		FROM price_history
		WHERE game_id = ANY($1) AND lang = $2
		ORDER BY game_id, checked_at DESC`,
		pq.Array(gameIDs), lang,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []model.PriceHistory

	for rows.Next() {
		var row model.PriceHistory

		err := rows.Scan(
			&row.ID,
			&row.GameID,
			&row.Lang,
			&row.Price,
			&row.Currency,
			&row.DiscountPercent,
			&row.CheckedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
