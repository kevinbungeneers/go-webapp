package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	code         int
	bytesWritten int
}

func (w *wrappedWriter) Write(b []byte) (int, error) {
	if w.code == 0 {
		w.WriteHeader(http.StatusOK)
	}

	n, err := w.ResponseWriter.Write(b)
	w.bytesWritten += n

	return n, err
}

func (w *wrappedWriter) WriteHeader(code int) {
	if w.code == 0 {
		w.code = code
		w.ResponseWriter.WriteHeader(code)
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		start := time.Now()

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		ww := wrappedWriter{ResponseWriter: w}
		slog.InfoContext(ctx, "request started",
			slog.String("http_scheme", scheme),
			slog.String("http_proto", r.Proto),
			slog.String("http_method", r.Method),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
			slog.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)),
		)

		defer func() {
			duration := time.Since(start)

			var logFunc func(ctx context.Context, msg string, args ...any)
			if ww.code >= http.StatusInternalServerError {
				logFunc = slog.ErrorContext
			} else {
				logFunc = slog.InfoContext
			}

			logFunc(ctx, "request finished",
				slog.Int("status", ww.code),
				slog.Int("bytes_written", ww.bytesWritten),
				slog.String("duration", duration.String()),
			)
		}()

		next.ServeHTTP(&ww, r)
	})
}
