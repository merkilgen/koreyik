package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/serwennn/koreyik/internal/models"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"net/http"
	"strconv"
)

type mediaEntryImpl struct{}

func registerMediaEntry(r chi.Router, stg *pq.Storage) {
	impl := &mediaEntryImpl{}

	r.Get("/media/{id}", impl.getMediaEntries(stg))
	r.Post("/media/", impl.postMediaEntries(stg))
}

func (impl *mediaEntryImpl) getMediaEntries(stg *pq.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		anime, err := models.GetAnime(stg, r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		animeJson, _ := json.Marshal(anime)
		w.Write(animeJson)
	}
}

func (impl *mediaEntryImpl) postMediaEntries(stg *pq.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create
	}
}
