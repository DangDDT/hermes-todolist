package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type traceKey string

const (
	// TraceIDKey is the context key for trace ID.
	TraceIDKey traceKey = "trace_id"
	// TraceIDHeader is the HTTP header name for trace ID.
	TraceIDHeader = "X-Trace-ID"
)

// GetTraceID extracts trace ID from context.
func GetTraceID(ctx context.Context) string {
	if id, ok := ctx.Value(TraceIDKey).(string); ok {
		return id
	}
	return ""
}

// generateTraceID creates a random 16-byte hex string.
func generateTraceID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("fallback-%d", middleware.NextRequestID())
	}
	return hex.EncodeToString(b)
}

// Tracing returns middleware that injects a trace ID into the request context
// and response header. If the client sends X-Trace-ID, it is forwarded.
func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use client-provided trace ID or generate one
		traceID := r.Header.Get(TraceIDHeader)
		if traceID == "" {
			traceID = generateTraceID()
		}

		// Store in context
		ctx := context.WithValue(r.Context(), TraceIDKey, traceID)
		r = r.WithContext(ctx)

		// Set response header
		w.Header().Set(TraceIDHeader, traceID)

		next.ServeHTTP(w, r)
	})
}
