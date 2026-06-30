package auth_register

import (
	"context"

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
	// Check if username already exists.
	existing, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	if existing != nil {
		return nil, apperrors.Conflict("User", "username", nil)
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
