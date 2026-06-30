package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/user"
)

// UserRepo implements user.UserRepository using PostgreSQL.
type UserRepo struct {
	pool *pgxpool.Pool
}

// NewUserRepo creates a new UserRepo.
func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

// Create persists a new user.
func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	return ErrNotImplemented
}

// GetByUsername retrieves a user by username.
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	return nil, ErrNotImplemented
}

// GetByID retrieves a user by ID.
func (r *UserRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	return nil, ErrNotImplemented
}
