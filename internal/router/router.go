package router

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	gamehandler "go.rest.api/internal/handler/game"
	trackedgamehandler "go.rest.api/internal/handler/tracked_game"
	"go.rest.api/internal/httputil"
	"go.rest.api/internal/middleware"
	userrepo "go.rest.api/internal/repository/user"
)

func New(
	userRepo userrepo.UserRepository,
	gameHandler *gamehandler.Handler,
	trackedGameHandler *trackedgamehandler.Handler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Auth(userRepo))

		r.Get("/games", httputil.Wrap(gameHandler.Search))

		r.Route("/user/games", func(r chi.Router) {
			r.Post("/", httputil.Wrap(trackedGameHandler.Create))
			r.Patch("/{id}", httputil.Wrap(trackedGameHandler.Update))
			r.Delete("/{id}", httputil.Wrap(trackedGameHandler.Delete))
		})
	})

	return r
}
