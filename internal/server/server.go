package server

import (
	"github.com/serwennn/koreyik/internal/config"
	"net/http"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		HttpServer: &http.Server{
			Addr:        cfg.Address,
			Handler:     handler,
			ReadTimeout: cfg.Server.Timeout,
			IdleTimeout: cfg.Server.IdleTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.HttpServer.Shutdown(nil)
}
