package auth_login

import (
	"net/http"
	"time"

	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Handler handles POST /auth/login.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new login handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Login godoc
// @Summary      Login
// @Description  Authenticate user and return JWT token via cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginRequest  true  "Login payload"
// @Success      200      {object}  response.Envelope
// @Failure      400      {object}  response.Envelope
// @Failure      401      {object}  response.Envelope
// @Router       /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	req, err := DecodeLoginRequest(r)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid request body", err))
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, apperrors.ValidationError(err.Error(), err))
		return
	}

	result, err := h.usecase.Login(r.Context(), req)
	if err != nil {
		response.Error(w, err)
		return
	}

	// Set access token cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(result.ExpiresAt, 0),
	})

	response.Success(w, http.StatusOK, result, nil)
}
