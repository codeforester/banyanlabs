package logging

import (
	"io"
	"log/slog"
	"strings"
	"time"
)

func New(writer io.Writer, levelName string, service string) *slog.Logger {
	level := parseLevel(levelName)
	options := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			switch attr.Key {
			case slog.TimeKey:
				attr.Key = "timestamp"
				if timestamp, ok := attr.Value.Any().(time.Time); ok {
					attr.Value = slog.StringValue(timestamp.UTC().Format(time.RFC3339Nano))
				}
			case slog.LevelKey:
				attr.Value = slog.StringValue(strings.ToLower(attr.Value.String()))
			case slog.MessageKey:
				attr.Key = "message"
			}
			return attr
		},
	}

	return slog.New(slog.NewJSONHandler(writer, options)).With(slog.String("service", service))
}

func parseLevel(levelName string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(levelName)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
