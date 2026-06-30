package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config holds all configuration for the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
	CORS     CORSConfig
	Sentry   SentryConfig
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	Port int `env:"SERVER_PORT" envDefault:"8080"`
}

// DatabaseConfig holds database connection configuration.
type DatabaseConfig struct {
	URL string `env:"DATABASE_URL,required"`
}

// JWTConfig holds JWT authentication configuration.
type JWTConfig struct {
	Secret      string        `env:"JWT_SECRET,required"`
	ExpiryHours time.Duration `env:"JWT_EXPIRY_HOURS" envDefault:"168h"` // 7 days
}

// LogConfig holds logging configuration.
type LogConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"info"`
}

// CORSConfig holds CORS configuration.
type CORSConfig struct {
	Origin string `env:"CORS_ORIGIN" envDefault:"http://localhost:3000"`
}

// SentryConfig holds Sentry error tracking configuration.
type SentryConfig struct {
	DSN string `env:"SENTRY_DSN"`
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}

// MustLoad reads configuration or panics (for main.go).
func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}
