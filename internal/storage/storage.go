package storage

import (
	"fmt"
	"github.com/serwennn/koreyik/internal/config"
)

func DatabaseUrlCreator(storage config.Storage) string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}
