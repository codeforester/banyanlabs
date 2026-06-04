package sqlite

import (
	"context"
	"database/sql"
	"testing"
)

func TestRunMigrationsCreatesInitialSchema(t *testing.T) {
	ctx := context.Background()
	db, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	defer db.Close()

	if err := RunMigrations(ctx, db); err != nil {
		t.Fatalf("run migrations: %v", err)
	}
	if err := RunMigrations(ctx, db); err != nil {
		t.Fatalf("run migrations second time: %v", err)
	}

	for _, table := range []string{"users", "sessions", "shortened_urls", "short_codes", "schema_migrations"} {
		if !tableExists(ctx, t, db, table) {
			t.Fatalf("expected table %q to exist", table)
		}
	}
	if !migrationRecorded(ctx, t, db, "001_initial") {
		t.Fatal("expected migration 001_initial to be recorded")
	}
}

func tableExists(ctx context.Context, t *testing.T, db *sql.DB, table string) bool {
	t.Helper()

	var name string
	err := db.QueryRowContext(
		ctx,
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = ?",
		table,
	).Scan(&name)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		t.Fatalf("query table %q: %v", table, err)
	}
	return name == table
}

func migrationRecorded(ctx context.Context, t *testing.T, db *sql.DB, version string) bool {
	t.Helper()

	var found string
	err := db.QueryRowContext(ctx, "SELECT version FROM schema_migrations WHERE version = ?", version).Scan(&found)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		t.Fatalf("query migration %q: %v", version, err)
	}
	return found == version
}
