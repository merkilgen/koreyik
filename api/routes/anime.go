package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/serwennn/koreyik/internal/models"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"github.com/serwennn/koreyik/internal/templates"
	"log/slog"
	"net/http"
	"strconv"
)

type animeImpl struct{}

func registerAnime(r chi.Router, stg *pq.Storage, log *slog.Logger) {
	impl := &animeImpl{}

	r.Route("/anime", func(r chi.Router) {
		r.Get("/{id}", impl.getAnime(stg, log))
		r.Post("/", impl.postAnime(stg))
	})
}

func (impl *animeImpl) getAnime(stg *pq.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		/* UNCOMMENT THIS ONE DAY

		// Try to get an entry from Redis
		pair := fmt.Sprintf("anime:%d", id)
		cached, err := cacheServer.Client.JSONGet(r.Context(), pair, "$").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// If Redis has the entry, then return it
		if cached != "" {
			w.Write([]byte(cached))
			log.Debug("Got an entry from the cache")
			return
		}

		*/

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

		// Serialize go struct to show it in json format
		_, err = json.Marshal(anime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		/* UNCOMMENT THIS ONE DAY

		// Set entry in Redis
		err = cacheServer.Client.JSONSet(r.Context(), pair, "$", serialized).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = cacheServer.Client.Expire(r.Context(), pair, 30*time.Second).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debug("Wrote an entry to the cache", slog.Int("ttl", 30))

		*/

		err = templates.Template.ExecuteTemplate(w, "anime.html", anime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
