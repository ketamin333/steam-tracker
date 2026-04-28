package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"go.rest.api/internal/apperr"
	"go.rest.api/internal/httputil"
	userrepo "go.rest.api/internal/repository/user"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func Auth(repo userrepo.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				httputil.Error(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				httputil.Error(w, http.StatusUnauthorized, "invalid authorization format")
				return
			}

			apiKey := strings.TrimSpace(parts[1])
			user, err := repo.GetByAPIKey(r.Context(), apiKey)

			if errors.Is(err, apperr.ErrNotFound) {
				httputil.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if err != nil {
				httputil.Error(w, http.StatusInternalServerError, "server error")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
