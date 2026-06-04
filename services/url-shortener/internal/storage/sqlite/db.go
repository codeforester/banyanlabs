package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Open(ctx context.Context, databasePath string) (*sql.DB, error) {
	if databasePath == "" {
		return nil, fmt.Errorf("database path is required")
	}
	if err := ensureDatabaseDirectory(databasePath); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dsn(databasePath))
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		closeErr := db.Close()
		if closeErr != nil {
			return nil, fmt.Errorf("ping database: %w; close database: %v", err, closeErr)
		}
		return nil, err
	}
	return db, nil
}

func dsn(databasePath string) string {
	if databasePath == ":memory:" {
		return "file::memory:?cache=shared&_foreign_keys=on&_busy_timeout=5000"
	}
	return fmt.Sprintf("file:%s?_foreign_keys=on&_busy_timeout=5000", databasePath)
}

func ensureDatabaseDirectory(databasePath string) error {
	if databasePath == ":memory:" {
		return nil
	}
	directory := filepath.Dir(databasePath)
	if directory == "." || directory == "" {
		return nil
	}
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create database directory %q: %w", directory, err)
	}
	return nil
}
