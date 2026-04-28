package trackedgamerepo

import (
	"context"

	"go.rest.api/internal/apperr"
)

func (r *Repository) Remove(ctx context.Context, userID, gameID int) error {
	result, err := r.db.ExecContext(ctx,
		`DELETE FROM tracked_games WHERE user_id = $1 AND game_id = $2`,
		userID, gameID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return apperr.ErrNotFound
	}

	return nil
}
