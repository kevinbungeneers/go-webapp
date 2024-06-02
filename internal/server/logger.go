package server

import (
	"context"
	"go-webapp/internal"
	"io"
	"log/slog"
)

type defaultHandler struct {
	handler slog.Handler
}

func (d defaultHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return d.handler.Enabled(ctx, level)
}

func (d defaultHandler) Handle(ctx context.Context, record slog.Record) error {
	requestId, ok := internal.RequestIdFromContext(ctx)
	if ok {
		record.AddAttrs(slog.String("request_id", requestId))
	}

	return d.handler.Handle(ctx, record)
}

func (d defaultHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d.handler.WithAttrs(attrs)
}

func (d defaultHandler) WithGroup(name string) slog.Handler {
	return d.handler.WithGroup(name)
}

func configureLogger(w io.Writer, opts *slog.HandlerOptions) {
	logger := slog.New(defaultHandler{handler: slog.NewTextHandler(w, opts)})
	slog.SetDefault(logger)
}
