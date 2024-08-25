package pq

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5" // driver
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/storage"
)

type Storage struct {
	conn *pgx.Conn
}

func New(storageOptions config.Storage) (*Storage, error) {
	url := storage.DatabaseUrlCreator(storageOptions)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer conn.Close(context.Background())

	// Check the connection
	if err = conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &Storage{conn: conn}, nil
}
