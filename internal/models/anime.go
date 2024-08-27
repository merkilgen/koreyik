package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
)

// Anime

type Anime struct {
	MediaEntry
	//Info
}

/*
type Info struct {
	Seasons  int `json:"seasons"`
	Episodes int `json:"episodes"`

	Genres []string `json:"genres"`
	Themes []string `json:"themes"`

	Producers []string `json:"producers"`
	Studios   []string `json:"studios"`
}
*/

func CreateAnime(storage *pq.Storage, ctx context.Context, id int, thumbnailUrl, titleKk, titleEn, titleJp string) (Anime, error) {
	_, err := storage.Exec(ctx, "INSERT INTO animes VALUES($1, $2, $3, $4, $5)", id, thumbnailUrl, titleKk, titleEn, titleJp)
	if err != nil {
		return Anime{}, err
	}
	return Anime{}, nil
}

func GetAnime(storage *pq.Storage, ctx context.Context, id int) (Anime, error) {
	row, _ := storage.Query(ctx, "SELECT * FROM animes WHERE id = $1", id)

	anime, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Anime])
	if err != nil {
		return Anime{}, err
	}
	return anime, nil
}
