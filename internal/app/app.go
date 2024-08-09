package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg := config.New()

	// stg := storage.New(cfg.StoragePath)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Koreyik!"))
	})

	srv := server.NewServer(cfg, r)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Failed to run the server: %s\n", err.Error())
		}
	}()

	log.Printf("Server is running on http://%s [ENV: %s]\n", srv.HttpServer.Addr, cfg.Env)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// TODO: Close the storage

	log.Println("Shutting down the server...")
}
