package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage"
	"log"
	"net/http"
)

func main() {
	cfg := config.New()

	stg := storage.New(cfg.StoragePath)
	_ = stg // temporary

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Koreyik!"))
	})

	// TODO: Graceful shutdown

	serv := server.NewServer(cfg, r)
	log.Printf("Server is running on http://%s [ENV: %s]\n", serv.HttpServer.Addr, cfg.Env)
	log.Fatal(serv.Run())
}
