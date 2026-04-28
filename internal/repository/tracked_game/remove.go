package trackedgamerepo

import "context"

func (r *Repository) Remove(ctx context.Context, userID, gameID int) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM tracked_games WHERE user_id = $1 AND game_id = $2`,
		userID, gameID,
	)

	return err
}
