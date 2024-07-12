package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"koreyik/api/handler"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handler.SetupRoutes(r)

	log.Print("Listening server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
