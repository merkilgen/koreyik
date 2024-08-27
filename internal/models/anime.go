package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"time"
)

type Anime struct {
	ID           int
	ThumbnailURL string
	Description  string
	Rating       string

	TitleKk string
	TitleJp string
	TitleEn string

	Status         string
	StartedAiring  time.Time
	FinishedAiring time.Time

	Genres []string
	Themes []string

	Seasons  int
	Episodes int
	Duration int

	Studios   []string
	Producers []string

	//Related []MediaEntry
}

func CreateAnime(storage *pq.Storage, ctx context.Context, a Anime) error {
	fmt.Println(a)
	fmt.Println(fmt.Sprintf("%T", a))

	an := Anime{
		ID:             3,
		ThumbnailURL:   "2",
		Description:    "a.Descriipton",
		Rating:         "a.Rating",
		TitleKk:        "a.TitleKk",
		TitleJp:        "a.TitleJp",
		TitleEn:        "a.TitleEn",
		Status:         "a.Status",
		StartedAiring:  time.Now(),
		FinishedAiring: time.Now(),
		Genres:         []string{"a.Genres"},
		Themes:         []string{"a.Genres"},
		Seasons:        1,
		Episodes:       2,
		Duration:       3,
		Studios:        []string{"a.Genres"},
		Producers:      []string{"a.Genres"},
	}
	_, err := storage.Exec(ctx, "INSERT INTO animes VALUES (?)", an)
	if err != nil {
		return err
	}
	return nil
}

func GetAnime(storage *pq.Storage, ctx context.Context, id int) (Anime, error) {
	row, _ := storage.Query(ctx, "SELECT * FROM animes WHERE id = $1", id)

	anime, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Anime])
	if err != nil {
		return Anime{}, err
	}
	return anime, nil
}
