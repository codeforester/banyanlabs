package config

import "os"

const (
	defaultAddr         = "127.0.0.1:8080"
	defaultDatabasePath = "var/url-shortener/url-shortener.sqlite3"
	defaultLogLevel     = "info"
	serviceName         = "url-shortener"
)

type Config struct {
	Addr         string
	DatabasePath string
	LogLevel     string
	ServiceName  string
}

func Load() Config {
	return LoadFromEnv(os.Getenv)
}

func LoadFromEnv(getenv func(string) string) Config {
	return Config{
		Addr:         valueOrDefault(getenv("BANYAN_URL_SHORTENER_ADDR"), defaultAddr),
		DatabasePath: valueOrDefault(getenv("BANYAN_URL_SHORTENER_DATABASE"), defaultDatabasePath),
		LogLevel:     valueOrDefault(getenv("BANYAN_LOG_LEVEL"), defaultLogLevel),
		ServiceName:  serviceName,
	}
}

func valueOrDefault(value string, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
