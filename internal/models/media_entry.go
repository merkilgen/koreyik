package models

import "time"

type MediaEntry struct {
	ID           int          `json:"id"`
	ThumbnailURL string       `json:"thumbnail_url"`
	Titles       Titles       `json:"titles"`
	Related      []MediaEntry `json:"related"`

	Status         string    `json:"status"`
	StartedAiring  time.Time `json:"started_airing"`
	FinishedAiring time.Time `json:"finished_airing"`

	Rating string `json:"rating"`
}

type Titles struct {
	Kk string `json:"kk"`
	Jp string `json:"jp"`
	En string `json:"en"`
}
