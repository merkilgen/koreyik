package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"log/slog"
)

func RegisterRoutes(r *chi.Mux, stg *pq.Storage, log *slog.Logger) {
	// Register the API routes
	api := chi.NewRouter()

	registerAnime(api, stg, log)

	r.Mount("/api", api)
}
