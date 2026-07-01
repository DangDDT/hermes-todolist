package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

type contextKey string

const userIDKey contextKey = "user_id"

// Skip paths that don't require authentication.
var skipPaths = map[string]bool{
	"/api/v1/health":        true,
	"/api/v1/auth/login":    true,
	"/api/v1/auth/register": true,
	"/swagger":              true,
}

// JWTAuth returns a Chi middleware that validates JWT tokens from cookies.
func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for certain paths.
			if shouldSkipAuth(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			// Extract token from cookie.
			cookie, err := r.Cookie("access_token")
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, map[string]string{
					"error": "missing access token",
				})
				return
			}

			userID, err := validateToken(cookie.Value, secret)
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID extracts the authenticated user's ID from the context.
func GetUserID(ctx context.Context) (uuid.UUID, error) {
	id, ok := ctx.Value(userIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID not found in context")
	}
	return id, nil
}

// shouldSkipAuth returns true if the path doesn't require authentication.
func shouldSkipAuth(path string) bool {
	for skip := range skipPaths {
		if len(path) >= len(skip) && path[:len(skip)] == skip {
			return true
		}
	}
	return false
}

// validateToken parses and validates a JWT token, returning the user ID.
func validateToken(tokenString, secret string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	// Check expiration.
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return uuid.Nil, fmt.Errorf("token expired")
		}
	}

	// Extract user_id.
	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user_id in token: %w", err)
	}

	return userID, nil
}
