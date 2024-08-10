package server

import (
	"github.com/serwennn/koreyik/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        cfg.Address,
			Handler:     handler,
			ReadTimeout: cfg.Server.Timeout,
			IdleTimeout: cfg.Server.IdleTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.httpServer.Shutdown(nil)
}
