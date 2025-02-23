package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// ANSI color codes for log levels
var levelColors = map[slog.Level]string{
	slog.LevelDebug: "\033[36m", // Cyan
	slog.LevelInfo:  "\033[32m", // Green
	slog.LevelWarn:  "\033[33m", // Yellow
	slog.LevelError: "\033[31m", // Red
}

const resetColor = "\033[0m" // Reset color

// Custom text handler with color support
type ColorTextHandler struct {
	handler slog.Handler
}

func (h *ColorTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ColorTextHandler) Handle(ctx context.Context, r slog.Record) error {
	color, ok := levelColors[r.Level]
	if !ok {
		color = resetColor
	}

	// Format timestamp as ISO8601
	timestamp := time.Now().Format(time.RFC3339)

	// Custom log output format: [TIME] [LEVEL] MESSAGE key=value ...
	msg := fmt.Sprintf("%s[%s] %s[%s]%s %s",
		color, r.Level.String(), resetColor, timestamp, resetColor, r.Message)

	// Iterate over attributes (structured fields)
	r.Attrs(func(a slog.Attr) bool {
		msg += fmt.Sprintf(" %s=%v", a.Key, a.Value.Any())
		return true
	})

	fmt.Println(msg)
	return nil
}

func (h *ColorTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ColorTextHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *ColorTextHandler) WithGroup(name string) slog.Handler {
	return &ColorTextHandler{handler: h.handler.WithGroup(name)}
}

func Get(level slog.Level) *slog.Logger {
	return slog.New(&ColorTextHandler{
		handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	})
}
