package router

import (
	gamehandler "steam-tracker/internal/handler/game"
	trackedgamehandler "steam-tracker/internal/handler/tracked_game"
	userhandler "steam-tracker/internal/handler/user"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	userrepo "steam-tracker/internal/repository/user"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func New(
	userRepo userrepo.UserRepository,
	gameHandler *gamehandler.Handler,
	trackedGameHandler *trackedgamehandler.Handler,
	userHandler *userhandler.Handler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Auth(userRepo))

		r.Route("/user", func(r chi.Router) {
			r.Get("/", httputil.Wrap(userHandler.Get))
			r.Patch("/", httputil.Wrap(userHandler.Update))
		})

		r.Get("/games", httputil.Wrap(gameHandler.Search))

		r.Route("/user/games", func(r chi.Router) {
			r.Post("/", httputil.Wrap(trackedGameHandler.Create))
			r.Patch("/{id}", httputil.Wrap(trackedGameHandler.Update))
			r.Delete("/{id}", httputil.Wrap(trackedGameHandler.Delete))
		})
	})

	return r
}
