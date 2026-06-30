package auth_register

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/user"
)

type userRepoStub struct {
	getByUsername func(ctx context.Context, username string) (*user.User, error)
	create        func(ctx context.Context, u *user.User) error
	getByID       func(ctx context.Context, id uuid.UUID) (*user.User, error)
}

func (s userRepoStub) Create(ctx context.Context, u *user.User) error {
	if s.create != nil {
		return s.create(ctx, u)
	}
	return nil
}

func (s userRepoStub) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	if s.getByUsername != nil {
		return s.getByUsername(ctx, username)
	}
	return nil, nil
}

func (s userRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	if s.getByID != nil {
		return s.getByID(ctx, id)
	}
	return nil, nil
}

func TestRegisterReturnsConflictWhenUsernameExists(t *testing.T) {
	existing, err := user.NewUser("dang", "Dang")
	if err != nil {
		t.Fatal(err)
	}
	uc := NewUsecase(userRepoStub{
		getByUsername: func(ctx context.Context, username string) (*user.User, error) {
			return existing, nil
		},
		create: func(ctx context.Context, u *user.User) error {
			t.Fatal("create should not be called when username exists")
			return nil
		},
	})

	_, err = uc.Register(context.Background(), &RegisterUserRequest{Username: "dang", Password: "secret123", DisplayName: "Dang"})
	if err == nil || !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("expected conflict error, got %v", err)
	}
}

func TestRegisterPropagatesUsernameLookupFailure(t *testing.T) {
	lookupErr := errors.New("db unavailable")
	uc := NewUsecase(userRepoStub{
		getByUsername: func(ctx context.Context, username string) (*user.User, error) {
			return nil, lookupErr
		},
		create: func(ctx context.Context, u *user.User) error {
			t.Fatal("create should not be called when username lookup fails")
			return nil
		},
	})

	_, err := uc.Register(context.Background(), &RegisterUserRequest{Username: "dang", Password: "secret123", DisplayName: "Dang"})
	if err == nil || !strings.Contains(err.Error(), "internal") {
		t.Fatalf("expected internal error from lookup failure, got %v", err)
	}
}

func TestRegisterCreatesUserWithHashedPassword(t *testing.T) {
	var created *user.User
	uc := NewUsecase(userRepoStub{
		create: func(ctx context.Context, u *user.User) error {
			created = u
			return nil
		},
	})

	resp, err := uc.Register(context.Background(), &RegisterUserRequest{Username: "dang", Password: "secret123", DisplayName: "Dang"})
	if err != nil {
		t.Fatalf("register failed: %v", err)
	}
	if resp.Username != "dang" || resp.DisplayName != "Dang" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if created == nil {
		t.Fatal("expected user to be persisted")
	}
	if created.PasswordHash == "" || created.PasswordHash == "secret123" {
		t.Fatalf("expected password to be hashed, got %q", created.PasswordHash)
	}
	if err := created.CheckPassword("secret123"); err != nil {
		t.Fatalf("hashed password should verify: %v", err)
	}
}
