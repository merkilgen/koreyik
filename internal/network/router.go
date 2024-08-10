package network

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.Recoverer,
		middleware.Logger,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Koreyik!"))
	})

	return r
}
