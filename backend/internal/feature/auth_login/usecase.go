package auth_login

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/user"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
)

// Usecase handles user login.
type Usecase struct {
	userRepo   user.UserRepository
	jwtSecret  string
	expiryHours time.Duration
}

// NewUsecase creates a new login usecase.
func NewUsecase(userRepo user.UserRepository, jwtSecret string, expiryHours time.Duration) *Usecase {
	return &Usecase{
		userRepo:   userRepo,
		jwtSecret:  jwtSecret,
		expiryHours: expiryHours,
	}
}

// Login authenticates a user and returns a JWT token.
func (uc *Usecase) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	u, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, apperrors.Unauthorized("invalid username or password", err)
	}
	if u == nil {
		return nil, apperrors.Unauthorized("invalid username or password", nil)
	}

	if err := u.CheckPassword(req.Password); err != nil {
		return nil, apperrors.Unauthorized("invalid username or password", err)
	}

	expiresAt := time.Now().Add(uc.expiryHours)
	token, err := uc.generateJWT(u, expiresAt)
	if err != nil {
		return nil, apperrors.Internal(fmt.Errorf("failed to generate token: %w", err))
	}

	return &LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt.Unix(),
		User: UserInfo{
			ID:          u.ID.String(),
			Username:    u.Username,
			DisplayName: u.DisplayName,
			CreatedAt:   u.CreatedAt,
		},
	}, nil
}

func (uc *Usecase) generateJWT(u *user.User, expiresAt time.Time) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  u.ID.String(),
		"username": u.Username,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.jwtSecret))
}
