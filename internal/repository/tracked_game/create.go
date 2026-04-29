package trackedgamerepo

import (
	"context"
	"database/sql"
	"errors"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/model"
)

func (r *Repository) Create(ctx context.Context, user *model.User, gameID int, price *float64) (*model.TrackedGame, error) {
	var g model.TrackedGame

	err := r.db.QueryRowContext(ctx,
		`INSERT INTO tracked_games (user_id, game_id, target_price) VALUES ($1, $2, $3)
			ON CONFLICT (user_id, game_id) DO NOTHING
			RETURNING id, user_id, game_id, target_price, created_at`,
		user.ID, gameID, price,
	).Scan(&g.ID, &g.UserID, &g.GameID, &g.TargetPrice, &g.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperr.ErrAlreadyExists
	}

	if err != nil {
		return nil, err
	}

	return &g, nil
}
