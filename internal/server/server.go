package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	port   string
	router *chi.Mux
}

func NewServer(p string, r *chi.Mux) *Server {
	return &Server{port: p, router: r}
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.port, s.router)
}
