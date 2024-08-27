package pq

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func databaseUrlCreator(storage config.Storage) string {
	// URL should look like this -> "postgres://username:password@host:port/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}

func (s *Storage) Shutdown() {
	s.db.Close()
}

func (s *Storage) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return s.db.Exec(ctx, query, args...)
}

func (s *Storage) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return s.db.Query(ctx, query, args...)
}

func (s *Storage) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return s.db.QueryRow(ctx, query, args...)
}
