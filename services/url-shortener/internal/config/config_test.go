package config

import "testing"

func TestLoadFromEnvUsesDefaults(t *testing.T) {
	cfg := LoadFromEnv(func(string) string { return "" })

	if cfg.Addr != defaultAddr {
		t.Fatalf("Addr = %q, want %q", cfg.Addr, defaultAddr)
	}
	if cfg.DatabasePath != defaultDatabasePath {
		t.Fatalf("DatabasePath = %q, want %q", cfg.DatabasePath, defaultDatabasePath)
	}
	if cfg.LogLevel != defaultLogLevel {
		t.Fatalf("LogLevel = %q, want %q", cfg.LogLevel, defaultLogLevel)
	}
	if cfg.ServiceName != serviceName {
		t.Fatalf("ServiceName = %q, want %q", cfg.ServiceName, serviceName)
	}
}

func TestLoadFromEnvUsesOverrides(t *testing.T) {
	values := map[string]string{
		"BANYAN_URL_SHORTENER_ADDR":     "127.0.0.1:9090",
		"BANYAN_URL_SHORTENER_DATABASE": "tmp/test.sqlite3",
		"BANYAN_LOG_LEVEL":              "debug",
	}

	cfg := LoadFromEnv(func(key string) string { return values[key] })

	if cfg.Addr != "127.0.0.1:9090" {
		t.Fatalf("Addr = %q", cfg.Addr)
	}
	if cfg.DatabasePath != "tmp/test.sqlite3" {
		t.Fatalf("DatabasePath = %q", cfg.DatabasePath)
	}
	if cfg.LogLevel != "debug" {
		t.Fatalf("LogLevel = %q", cfg.LogLevel)
	}
}
