package pq

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serwennn/koreyik/internal/config"
)

type Storage struct {
	DB *pgxpool.Pool
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

	return &Storage{DB: db}, nil
}

func databaseUrlCreator(storage config.Storage) string {
	// URL should look like this -> "postgres://username:password@host:port/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}

func (s *Storage) Shutdown() {
	s.DB.Close()
}
