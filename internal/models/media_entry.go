package models

import (
	"database/sql"
)

type MediaEntry struct {
	ID           int
	ThumbnailURL sql.NullString

	TitleKk sql.NullString
	TitleJp sql.NullString
	TitleEn sql.NullString

	//Related []MediaEntry
	//
	//Status         string
	//StartedAiring  time.Time
	//FinishedAiring time.Time
	//
	//Rating string
}
