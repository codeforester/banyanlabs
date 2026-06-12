package app

import (
	"context"
	"testing"
	"time"
)

func TestHealthReportsServiceStatusAndUptime(t *testing.T) {
	application := New(Options{})

	time.Sleep(2 * time.Millisecond)
	health := application.Health(context.Background())

	if health.Service != "url-shortener" {
		t.Fatalf("service = %q, want %q", health.Service, "url-shortener")
	}
	if health.Status != "ok" {
		t.Fatalf("status = %q, want %q", health.Status, "ok")
	}

	uptime, err := time.ParseDuration(health.Uptime)
	if err != nil {
		t.Fatalf("parse uptime %q: %v", health.Uptime, err)
	}
	if uptime <= 0 {
		t.Fatalf("uptime = %s, want positive duration", health.Uptime)
	}
}
