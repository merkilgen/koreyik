package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/serwennn/koreyik/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.New()
	_ = cfg

	// TODO: Connect to database

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Koreyik!"))
	})

	serv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	// TODO: Graceful shutdown

	log.Fatal(serv.ListenAndServe())
}
