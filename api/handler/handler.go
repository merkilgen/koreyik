// api/handler/handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func SetupRoutes(r chi.Router) {
	r.Get("/", RootHandler)
}
