package game_repository

import (
	"context"

	"go.rest.api/internal/models"
)

func (r *GameRepository) Upsert(ctx context.Context, g *models.Game) error {
	return r.db.QueryRowContext(ctx,
		`INSERT INTO games (steam_app_id, name, cover_url)
		VALUES ($1, $2, $3)
		ON CONFLICT (steam_app_id) DO UPDATE
		SET name = EXCLUDED.name, 
		    cover_url = EXCLUDED.cover_url,
		    updated_at = NOW()
		RETURNING id, steam_app_id, name, cover_url, created_at, updated_at`,
		g.SteamAppID, g.Name, g.CoverURL,
	).Scan(&g.ID, &g.SteamAppID, &g.Name, &g.CoverURL, &g.CreatedAt, &g.UpdatedAt)
}
