package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/serwennn/koreyik/internal/models"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"github.com/serwennn/koreyik/internal/storage/red"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type animeImpl struct{}

func registerAnime(r chi.Router, stg *pq.Storage, cacheServer *red.CacheServer, log *slog.Logger) {
	impl := &animeImpl{}

	r.Get("/anime/{id}", impl.getAnime(stg, cacheServer, log))
	r.Post("/anime/", impl.postAnime(stg))
}

func (impl *animeImpl) getAnime(stg *pq.Storage, cacheServer *red.CacheServer, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Try to get a value from Redis
		key := "anime"
		cached, err := cacheServer.Get(r.Context(), key)
		if err != nil && !errors.Is(err, redis.Nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// If Redis has the value, then return it
		if !errors.Is(err, redis.Nil) {
			w.Write([]byte(cached))
			log.Debug("Got an entry from the cache")
			return
		}

		// Get an entry from the main storage
		anime, err := models.GetAnime(stg, r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		log.Debug("Got an entry from the storage")

		// Serialize go struct to store/show it in json format
		serialized, err := json.Marshal(anime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set key-value in Redis
		err = cacheServer.Set(r.Context(), key, serialized, 30*time.Second)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debug("Wrote an entry to the cache")

		w.Write(serialized)
		return
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
		http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(newAnime.ID)), http.StatusFound)
	}
}
