package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // driver for SQLite3
	"log"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) *Storage {
	// Connect to the database
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	} else {
		log.Println("Connected to the database")
	}

	// TODO: Create tables

	return &Storage{db: db}
}
