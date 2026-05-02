package userrepo

import (
	"context"
	"steam-tracker/internal/model"
)

func (r *Repository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (api_key, lang, email) 
		VALUES ($1, $2, $3)
		RETURNING id, api_key, lang, email, created_at`,
		user.APIKey, user.Lang, *user.Email,
	).Scan(&user.ID, &user.APIKey, &user.Lang, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
