package trackedgamerepo

import (
	"context"
	"steam-tracker/internal/model"
)

type TrackedGameRow struct {
	ID          int
	User        model.User
	Game        model.Game
	TargetPrice *float64
}

func (r *Repository) GetAll(ctx context.Context) ([]TrackedGameRow, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT 
    		tracked_games.id, 
       		users.id, users.lang, users.email, users.created_at,
       		games.id, games.steam_app_id, games.name, games.cover_url, games.created_at, games.updated_at,
       		tracked_games.target_price
        FROM tracked_games
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

		err := rows.Scan(
			&row.ID,
			&row.User.ID,
			&row.User.Lang,
			&row.User.Email,
			&row.User.CreatedAt,
			&row.Game.ID,
			&row.Game.SteamAppID,
			&row.Game.Name,
			&row.Game.CoverURL,
			&row.Game.CreatedAt,
			&row.Game.UpdatedAt,
			&row.TargetPrice,
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
