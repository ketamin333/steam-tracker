package game_repository

import "go.rest.api/internal/models"

func (r *GameRepository) GetAll() ([]models.Game, error) {
	rows, err := r.db.Query(
		`SELECT id, name, description, genre, platform, price, rating, cover_url, created_at, updated_at FROM games
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	games := make([]models.Game, 0)

	for rows.Next() {
		var g models.Game
		err := rows.Scan(
			&g.ID, &g.Name, &g.Description, &g.Genre,
			&g.Platform, &g.Price, &g.Rating, &g.CoverURL,
			&g.CreatedAt, &g.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}

	return games, nil
}
