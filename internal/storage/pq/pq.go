package pq

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5" // driver
	"github.com/serwennn/koreyik/internal/config"
)

type Storage struct {
	conn *pgx.Conn
}

func New(storageOptions config.Storage) (*Storage, error) {
	url := databaseUrlCreator(storageOptions)

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

func databaseUrlCreator(storage config.Storage) string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}
