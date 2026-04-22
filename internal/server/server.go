package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.rest.api/internal/config"
)

type Server struct {
	cfg    *config.Config
	router *chi.Mux
}

func NewServer(cfg *config.Config, r *chi.Mux) *Server {
	return &Server{cfg: cfg, router: r}
}

func (s *Server) Run() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", s.cfg.ServerPort),
		s.router,
	)
}
