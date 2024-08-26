package pq

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serwennn/koreyik/internal/config"
)

type Storage struct {
	db *pgxpool.Pool
}

var ctx = context.Background()

func New(storageConfig config.Storage) (*Storage, error) {
	url := databaseUrlCreator(storageConfig)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Shutdown() {
	s.db.Close()
}

func databaseUrlCreator(storage config.Storage) string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}
