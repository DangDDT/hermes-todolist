package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PoolConfig defines the connection pool configuration.
type PoolConfig struct {
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	HealthCheckPeriod time.Duration
}

// DefaultPoolConfig returns sensible defaults for the connection pool.
func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MaxConns:          25,
		MinConns:          5,
		MaxConnLifetime:   1 * time.Hour,
		HealthCheckPeriod: 1 * time.Minute,
	}
}

// NewPool creates a new pgxpool.Pool from a database URL.
func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	return NewPoolWithConfig(ctx, databaseURL, DefaultPoolConfig())
}

// NewPoolWithConfig creates a new pgxpool.Pool with custom configuration.
func NewPoolWithConfig(ctx context.Context, databaseURL string, cfg PoolConfig) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
	poolCfg.HealthCheckPeriod = cfg.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
