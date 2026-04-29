package trackedgamerepo

import (
	"context"
)

type TrackedGameRow struct {
	ID          int
	UserID      int
	GameID      int
	SteamAppID  int
	Lang        string
	TargetPrice *float64
}

func (r *Repository) GetAll(ctx context.Context) ([]TrackedGameRow, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT tracked_games.id, user_id, game_id, steam_app_id, lang, target_price FROM tracked_games
		INNER JOIN games ON tracked_games.game_id = games.id
		INNER JOIN users ON tracked_games.user_id = users.id`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []TrackedGameRow

	for rows.Next() {
		var row TrackedGameRow

		err := rows.Scan(&row.ID, &row.UserID, &row.GameID, &row.SteamAppID, &row.Lang, &row.TargetPrice)

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
