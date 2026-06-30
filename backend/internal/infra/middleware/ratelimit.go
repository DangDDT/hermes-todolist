package middleware

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter is a simple in-memory rate limiter.
type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new RateLimiter.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
	// Background cleanup every minute.
	go rl.cleanup()
	return rl
}

// RateLimit returns an HTTP middleware that enforces the rate limit.
func RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler {
	rl := NewRateLimiter(limit, window)
	return rl.Handler
}

// Handler is the Chi-compatible middleware handler.
func (rl *RateLimiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		rl.mu.Lock()
		now := time.Now()
		windowStart := now.Add(-rl.window)

		// Filter timestamps within the window.
		var recent []time.Time
		for _, t := range rl.requests[ip] {
			if t.After(windowStart) {
				recent = append(recent, t)
			}
		}

		if len(recent) >= rl.limit {
			rl.mu.Unlock()
			w.Header().Set("Retry-After", "60")
			http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
			return
		}

		recent = append(recent, now)
		rl.requests[ip] = recent
		rl.mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

// cleanup periodically removes expired entries.
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		cutoff := now.Add(-rl.window)
		for ip, times := range rl.requests {
			var recent []time.Time
			for _, t := range times {
				if t.After(cutoff) {
					recent = append(recent, t)
				}
			}
			if len(recent) == 0 {
				delete(rl.requests, ip)
			} else {
				rl.requests[ip] = recent
			}
		}
		rl.mu.Unlock()
	}
}
