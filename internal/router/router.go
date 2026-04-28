package router

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go.rest.api/internal/middleware"
	userrepo "go.rest.api/internal/repository/user"
)

func New(userRepo userrepo.UserRepository) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Auth(userRepo))
	})

	return r
}
