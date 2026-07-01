package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// Logger returns slog-based request logger middleware with tracing support.
// Logs request start (debug) and request completion with trace_id, status, duration.
func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Extract trace ID from context
			traceID := GetTraceID(r.Context())
			logger := logger.With("trace_id", traceID)

			// Get chi request ID if available
			reqID := middleware.GetReqID(r.Context())

			// Log request start at debug level
			attrs := []slog.Attr{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("query", r.URL.RawQuery),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			}
			if reqID != "" {
				attrs = append(attrs, slog.String("req_id", reqID))
			}
			logger.LogAttrs(r.Context(), slog.LevelDebug, "request started", attrs...)

			// Wrap response writer to capture status code and size
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Serve request (with recovery)
			defer func() {
				if rec := recover(); rec != nil {
					ww.WriteHeader(http.StatusInternalServerError)
					logger.LogAttrs(r.Context(), slog.LevelError, "request panicked",
						slog.Any("panic", rec),
						slog.String("method", r.Method),
						slog.String("path", r.URL.Path),
					)
				}
			}()
			next.ServeHTTP(ww, r)

			duration := time.Since(start)
			status := ww.Status()
			bytes := ww.BytesWritten()

			// Determine log level based on status
			level := slog.LevelInfo
			switch {
			case status >= 500:
				level = slog.LevelError
			case status >= 400:
				level = slog.LevelWarn
			}

			// Log request completion
			attrs = []slog.Attr{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("query", r.URL.RawQuery),
				slog.Int("status", status),
				slog.Int("bytes", bytes),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.Duration("duration", duration),
				slog.Float64("duration_ms", float64(duration.Microseconds())/1000.0),
			}
			if reqID != "" {
				attrs = append(attrs, slog.String("req_id", reqID))
			}
			logger.LogAttrs(r.Context(), level, "request completed", attrs...)
		})
	}
}
