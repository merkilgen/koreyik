package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // driver for SQLite3
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	// Connect to the database
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	// TODO: Create tables

	return &Storage{db: db}, nil
}
