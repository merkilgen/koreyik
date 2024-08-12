package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/network/routes"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func Run() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load .env file: %s", err.Error()))
	}

	cfg := config.New()

	log := setupLogger(cfg.Env)

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
			panic(fmt.Sprintf("Failed to run the server: %s", err.Error()))
		}
	}()

	log.Info(fmt.Sprintf("Server is running on http://%s", cfg.Address), "env", cfg.Env)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// TODO: Close the storage

	log.Info("Server is shutting down")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		panic(fmt.Sprintf("setupLogger: Env must be either %s or %s", EnvLocal, EnvProd))
	}

	return log
}
