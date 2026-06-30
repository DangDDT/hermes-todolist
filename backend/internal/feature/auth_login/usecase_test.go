package auth_login

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/user"
)

type userRepoStub struct {
	getByUsername func(ctx context.Context, username string) (*user.User, error)
	create        func(ctx context.Context, u *user.User) error
	getByID       func(ctx context.Context, id uuid.UUID) (*user.User, error)
}

func (s userRepoStub) Create(ctx context.Context, u *user.User) error {
	if s.create != nil { return s.create(ctx, u) }
	return nil
}
func (s userRepoStub) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	if s.getByUsername != nil { return s.getByUsername(ctx, username) }
	return nil, nil
}
func (s userRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	if s.getByID != nil { return s.getByID(ctx, id) }
	return nil, nil
}

func TestLoginReturnsSignedJWTForValidCredentials(t *testing.T) {
	u, err := user.NewUser("dang", "Dang")
	if err != nil { t.Fatal(err) }
	if err := u.SetPassword("secret123"); err != nil { t.Fatal(err) }
	secret := "test-secret"
	uc := NewUsecase(userRepoStub{getByUsername: func(ctx context.Context, username string) (*user.User, error) {
		if username != "dang" { t.Fatalf("unexpected username %q", username) }
		return u, nil
	}}, secret, time.Hour)

	resp, err := uc.Login(context.Background(), &LoginRequest{Username: "dang", Password: "secret123"})
	if err != nil { t.Fatalf("login failed: %v", err) }
	if resp.User.ID != u.ID.String() || resp.User.Username != "dang" { t.Fatalf("unexpected user info: %+v", resp.User) }

	parsed, err := jwt.Parse(resp.Token, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil || !parsed.Valid { t.Fatalf("expected valid jwt, got token=%v err=%v", parsed, err) }
	claims := parsed.Claims.(jwt.MapClaims)
	if claims["user_id"] != u.ID.String() || claims["username"] != "dang" { t.Fatalf("unexpected claims: %+v", claims) }
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	u, err := user.NewUser("dang", "Dang")
	if err != nil { t.Fatal(err) }
	if err := u.SetPassword("secret123"); err != nil { t.Fatal(err) }
	uc := NewUsecase(userRepoStub{getByUsername: func(ctx context.Context, username string) (*user.User, error) { return u, nil }}, "secret", time.Hour)

	_, err = uc.Login(context.Background(), &LoginRequest{Username: "dang", Password: "wrong"})
	if err == nil || !strings.Contains(err.Error(), "UNAUTHORIZED") { t.Fatalf("expected unauthorized error, got %v", err) }
}
