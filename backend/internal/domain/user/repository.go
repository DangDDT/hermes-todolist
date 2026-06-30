package user

import (
	"context"

	"github.com/google/uuid"
)

// UserRepository defines the persistence contract for users.
type UserRepository interface {
	Create(ctx context.Context, u *User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
}
