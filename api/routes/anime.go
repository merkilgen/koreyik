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

type animeImpl struct{}

func registerAnime(r chi.Router, stg *pq.Storage) {
	impl := &animeImpl{}

	r.Get("/anime/{id}", impl.getAnime(stg))
	r.Post("/anime/", impl.postAnime(stg))
}

func (impl *animeImpl) getAnime(stg *pq.Storage) http.HandlerFunc {
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

func (impl *animeImpl) postAnime(stg *pq.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAnime models.Anime

		err := json.NewDecoder(r.Body).Decode(&newAnime)
		if err != nil {
			http.Error(w, "Json decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = models.CreateAnime(stg, r.Context(), newAnime)
		if err != nil {
			http.Error(w, "Create Anime: "+err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		//http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(newAnime.ID)), http.StatusFound)
	}
}
