package game_repository

import (
	"database/sql"

	"go.rest.api/internal/models"
)

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) Create(g *models.Game) error {
	return r.db.QueryRow(
		`INSERT INTO games (name, description, genre, platform, price, rating, cover_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at;`,
		g.Name, g.Description, g.Genre, g.Platform, g.Price, g.Rating, g.CoverURL,
	).Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt)
}
