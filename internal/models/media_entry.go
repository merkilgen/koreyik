package models

// MediaEntry represents a media entry: anime, manga, light novel, etc.

type MediaEntry struct {
	ID           int          `json:"id"`
	ThumbnailURL string       `json:"thumbnail_url"`
	Info         Info         `json:"info"`
	Titles       Titles       `json:"titles"`
	Related      []MediaEntry `json:"related"`
}

type Info struct {
	Seasons  int `json:"seasons"`
	Episodes int `json:"episodes"`

	Status         Status `json:"status"`
	StartedAiring  string `json:"started_airing"`
	FinishedAiring string `json:"finished_airing"`

	Genres []string `json:"genres"`
	Themes []string `json:"themes"`

	Producers []string `json:"producers"`
	Studios   []string `json:"studios"`

	Rating string `json:"rating"`
}

type Titles struct {
	Kk string `json:"kk"`
	Jp string `json:"jp"`
	En string `json:"en"`
}

// Status represents the status of the media entry. It can be one of the following: "On Air", "Finished", "Upcoming".
type Status string

const (
	StatusOnAir    Status = "On Air"
	StatusFinished Status = "Finished"
	StatusUpcoming Status = "Upcoming"
)
