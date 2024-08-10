package app

import (
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/network"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg := config.New()

	stg := storage.New(cfg.StoragePath)
	_ = stg // TODO: Use the storage

	r := network.NewRouter()

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
