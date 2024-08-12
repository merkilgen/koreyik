package routes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	// Register the API routes
	r.Route("/api", func(r chi.Router) {
		registerMediaEntry(r)
	})
}
