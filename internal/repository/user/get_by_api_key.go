package userrepo

import (
	"context"
	"database/sql"
	"errors"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/model"
)

func (r *Repository) GetByAPIKey(ctx context.Context, apiKey string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, api_key, lang, created_at FROM users WHERE api_key = $1`,
		apiKey,
	).Scan(&user.ID, &user.APIKey, &user.Lang, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperr.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
