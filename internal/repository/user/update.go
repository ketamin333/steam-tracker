package userrepo

import (
	"context"
	"database/sql"
	"errors"
	"steam-tracker/internal/apperr"
	"steam-tracker/internal/model"
)

func (r *Repository) Update(ctx context.Context, user *model.User, email *string, lang *string) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		`UPDATE users SET 
        	email = COALESCE($1, email),
            lang = COALESCE($2, lang)
        WHERE id = $3
        RETURNING email, lang`,
		email, lang, user.ID,
	).Scan(&user.Email, &user.Lang)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperr.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
