package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	apperrors "go.rest.api/internal/errors"
	"go.rest.api/internal/handlers"
	userrepository "go.rest.api/internal/repository/user"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func Auth(repo *userrepository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				handlers.Error(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				handlers.Error(w, http.StatusUnauthorized, "invalid authorization format")
				return
			}

			apiKey := strings.TrimSpace(parts[1])

			user, err := repo.GetByAPIKey(r.Context(), apiKey)
			if errors.Is(err, apperrors.ErrNotFound) {
				handlers.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if err != nil {
				handlers.Error(w, http.StatusInternalServerError, "server error")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
