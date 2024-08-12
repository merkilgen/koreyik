package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/serwennn/koreyik/internal/models"
	"net/http"
	"strconv"
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

	mediaEntry := models.MediaEntry{
		ID: id,
		Titles: models.Titles{
			Kk: "Сынықтама",
			Jp: "テスト",
			En: "Test",
		},
		ThumbnailURL: "https://example.com/image.jpg",
		Info: models.Info{
			Seasons:        1,
			Episodes:       12,
			Status:         models.StatusFinished,
			StartedAiring:  "2021-01-01",
			FinishedAiring: "2021-03-01",
			Genres:         []string{"Action", "Adventure"},
			Themes:         []string{"Magic", "Fantasy"},
			Producers:      []string{"Aniplex", "Bones"},
			Studios:        []string{"Madhouse", "Sunrise"},
			Rating:         "PG-13",
		},
		Related: nil,
	}

	m, _ := json.Marshal(mediaEntry)
	w.Write(m)
}
