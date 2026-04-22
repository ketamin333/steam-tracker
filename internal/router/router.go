package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.rest.api/internal/handlers"
	gamehandler "go.rest.api/internal/handlers/game"
)

func Setup(gameHandler *gamehandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {

		r.Route("/games", func(r chi.Router) {
			r.Post("/", handlers.Wrap(gameHandler.Create))
			r.Get("/", handlers.Wrap(gameHandler.GetAll))
		})

	})

	return r
}
