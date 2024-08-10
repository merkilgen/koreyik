package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/network/routes"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()

	stg := storage.New(cfg.StoragePath)
	_ = stg // TODO: Use the storage

	// Router
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.Recoverer,
		middleware.Logger,
	)

	routes.RegisterRoutes(r)

	srv := server.NewServer(cfg, r)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Failed to run the server: %s\n", err.Error())
		}
	}()

	log.Printf("Server is running on http://%s [ENV: %s]\n", cfg.Address, cfg.Env)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// TODO: Close the storage

	log.Println("Shutting down the server...")
}
