package trackedgamerepo

import (
	"context"
)

func (r *Repository) Add(ctx context.Context, userID, gameID int) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO tracked_games (user_id, game_id) VALUES ($1, $2)
			ON CONFLICT (user_id, game_id) DO NOTHING`,
		userID, gameID,
	)

	return err
}
