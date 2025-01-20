package server

import (
	"entdemo/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(h *handler.Handler) (*Server, error) {
	return &Server{handler: h}, nil
}

func (s *Server) Run(port string) error {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// register routes
	s.handler.RegisterRoutes(r)

	if err := http.ListenAndServe(port, r); err != nil {
		return err
	}
	return nil
}
