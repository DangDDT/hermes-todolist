package auth_register

import (
	"net/http"

	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Handler handles POST /auth/register.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new register handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Register handles the registration request.
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	req, err := DecodeRegisterUserRequest(r)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid request body", err))
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, apperrors.ValidationError(err.Error(), err))
		return
	}

	result, err := h.usecase.Register(r.Context(), req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, result, nil)
}
