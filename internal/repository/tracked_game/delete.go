package trackedgamerepo

import (
	"context"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/model"
)

func (r *Repository) Delete(ctx context.Context, user *model.User, gameID int) error {
	result, err := r.db.ExecContext(ctx,
		`DELETE FROM tracked_games WHERE user_id = $1 AND game_id = $2`,
		user.ID, gameID,
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
