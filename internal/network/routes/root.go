package routes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	// Register the API routes
	api := chi.NewRouter()

	registerMediaEntry(api)

	r.Mount("/api", api)
}
