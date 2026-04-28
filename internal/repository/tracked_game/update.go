package trackedgamerepo

import (
	"context"
	"database/sql"
	"errors"

	"go.rest.api/internal/apperr"
	"go.rest.api/internal/model"
)

func (r *Repository) Update(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	var g model.TrackedGame

	err := r.db.QueryRowContext(ctx,
		`UPDATE tracked_games SET target_price = $1 
			WHERE user_id = $2 AND game_id = $3
			RETURNING id, user_id, game_id, target_price, created_at`,
		price, user.ID, gameID,
	).Scan(&g.ID, &g.UserID, &g.GameID, &g.TargetPrice, &g.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperr.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &g, nil
}
