package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRunReturnsErrorWhenServerCannotListen(t *testing.T) {
	t.Setenv("BANYAN_URL_SHORTENER_ADDR", "127.0.0.1:-1")
	t.Setenv("BANYAN_URL_SHORTENER_DATABASE", filepath.Join(t.TempDir(), "url-shortener.sqlite3"))
	t.Setenv("BANYAN_LOG_LEVEL", "error")

	err := run()
	if err == nil {
		t.Fatal("run() error = nil, want listen error")
	}
	if !strings.Contains(err.Error(), "listen") {
		t.Fatalf("run() error = %q, want listen error", err.Error())
	}
}
