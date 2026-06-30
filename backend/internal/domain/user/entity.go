package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	DisplayName  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new User with the given username and display name.
func NewUser(username, displayName string) (*User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	if displayName == "" {
		displayName = username
	}
	now := time.Now().UTC()
	return &User{
		ID:          uuid.New(),
		Username:    username,
		DisplayName: displayName,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// SetPassword hashes the plain-text password and stores it.
func (u *User) SetPassword(plain string) error {
	if plain == "" {
		return errors.New("password is required")
	}
	hash, err := HashPassword(plain)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	u.UpdatedAt = time.Now().UTC()
	return nil
}

// CheckPassword verifies the plain-text password against the stored hash.
func (u *User) CheckPassword(plain string) error {
	return VerifyPassword(u.PasswordHash, plain)
}
