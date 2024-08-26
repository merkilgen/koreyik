package models

// Anime

type Anime struct {
	MediaEntry
	Info
}

type Info struct {
	Seasons  int `json:"seasons"`
	Episodes int `json:"episodes"`

	Genres []string `json:"genres"`
	Themes []string `json:"themes"`

	Producers []string `json:"producers"`
	Studios   []string `json:"studios"`
}
