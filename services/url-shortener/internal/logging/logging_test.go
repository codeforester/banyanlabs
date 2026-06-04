package logging

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestLoggerUsesBanyanLabsSchemaKeys(t *testing.T) {
	var output bytes.Buffer
	logger := New(&output, "debug", "url-shortener")

	logger.Info("health checked", slog.String("component", "http"))

	var event map[string]any
	if err := json.Unmarshal(output.Bytes(), &event); err != nil {
		t.Fatalf("unmarshal log event: %v", err)
	}

	for _, key := range []string{"timestamp", "level", "message", "service", "component"} {
		if _, ok := event[key]; !ok {
			t.Fatalf("expected log key %q in %#v", key, event)
		}
	}
	if event["level"] != "info" {
		t.Fatalf("level = %v, want info", event["level"])
	}
	if event["message"] != "health checked" {
		t.Fatalf("message = %v", event["message"])
	}
	if event["service"] != "url-shortener" {
		t.Fatalf("service = %v", event["service"])
	}
}
