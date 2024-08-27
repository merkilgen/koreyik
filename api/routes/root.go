package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
)

func RegisterRoutes(r *chi.Mux, stg *pq.Storage) {
	// Register the API routes
	api := chi.NewRouter()

	registerMediaEntry(api, stg)

	r.Mount("/api", api)
}
