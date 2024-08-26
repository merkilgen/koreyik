package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/serwennn/koreyik/internal/models"
	"net/http"
	"strconv"
	"time"
)

type mediaEntryImpl struct{}

func registerMediaEntry(r chi.Router) {
	impl := &mediaEntryImpl{}

	r.Get("/media/{id}", impl.getMediaEntries)
}

func (impl *mediaEntryImpl) getMediaEntries(w http.ResponseWriter, r *http.Request) {
	// Here we should get the media entries from the storage by ID
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// For now, let's return a hardcoded media entry
	anime := models.Anime{
		MediaEntry: models.MediaEntry{
			ID:           id,
			ThumbnailURL: "https://example.com/image.jpg",
			Titles: models.Titles{
				Kk: "Сынықтама",
				Jp: "テスト",
				En: "Test",
			},
			Related: nil,

			Status:         "Finished",
			StartedAiring:  time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			FinishedAiring: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),

			Rating: "PG-13",
		},
		Info: models.Info{
			Seasons:  1,
			Episodes: 12,

			Genres: []string{"Action", "Adventure"},
			Themes: []string{"Magic", "Fantasy"},

			Producers: []string{"Aniplex", "Bones"},
			Studios:   []string{"Madhouse", "Sunrise"},
		},
	}

	m, _ := json.Marshal(anime)
	w.Write(m)
}
