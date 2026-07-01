package auth_register

import (
	"context"
	"log/slog"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/user"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
)

// Usecase handles user registration.
type Usecase struct {
	userRepo user.UserRepository
}

// NewUsecase creates a new register usecase.
func NewUsecase(userRepo user.UserRepository) *Usecase {
	return &Usecase{userRepo: userRepo}
}

// Register creates a new user account.
func (uc *Usecase) Register(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
// Check username exists.
	existing, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if existing != nil {
		return nil, apperrors.Conflict("User", "username", nil)
	}
	// If user doesn't exist (err may be ErrNotFound), proceed to create
	logger := slog.Default()
	if err != nil {
		logger.Warn("GetByUsername returned error during register (user may not exist)", "error", err)
	}

	// Create domain entity.
	u, err := user.NewUser(req.Username, req.DisplayName)
	if err != nil {
		return nil, apperrors.ValidationError(err.Error(), err)
	}

	if err := u.SetPassword(req.Password); err != nil {
		return nil, apperrors.Internal(err)
	}

	if err := uc.userRepo.Create(ctx, u); err != nil {
		return nil, apperrors.Internal(err)
	}

	return &RegisterUserResponse{
		ID:          u.ID.String(),
		Username:    u.Username,
		DisplayName: u.DisplayName,
		CreatedAt:   u.CreatedAt,
	}, nil
}
