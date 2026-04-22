package user_repository

import (
	"context"
	"database/sql"
	"errors"

	apperrors "go.rest.api/internal/errors"
	"go.rest.api/internal/models"
)

func (r *UserRepository) GetByAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	var user models.User

	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, api_key, created_at FROM users WHERE api_key = $1`,
		apiKey,
	).Scan(&user.ID, &user.APIKey, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperrors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
