package trackedgamerepo

import (
	"context"
	"database/sql"
	"errors"

	"go.rest.api/internal/apperr"
	"go.rest.api/internal/model"
)

func (r *Repository) Add(ctx context.Context, userID, gameID int) (*model.TrackedGame, error) {
	var g model.TrackedGame

	err := r.db.QueryRowContext(ctx,
		`INSERT INTO tracked_games (user_id, game_id) VALUES ($1, $2)
			ON CONFLICT (user_id, game_id) DO NOTHING
			RETURNING id, user_id, game_id, created_at`,
		userID, gameID,
	).Scan(&g.ID, &g.UserID, &g.GameID, &g.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperr.ErrAlreadyExists
	}

	if err != nil {
		return nil, err
	}

	return &g, nil
}
