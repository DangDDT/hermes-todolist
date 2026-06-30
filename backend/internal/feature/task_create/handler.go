package task_create

import (
	"net/http"

	"github.com/DangDDT/hermes-todolist/backend/internal/infra/middleware"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Handler handles POST /tasks.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new task create handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Create godoc
// @Summary      Create a task
// @Description  Create a new task (requires authentication)
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateTaskRequest  true  "Task payload"
// @Success      201      {object}  response.Envelope
// @Failure      400      {object}  response.Envelope
// @Failure      401      {object}  response.Envelope
// @Router       /tasks [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		response.Error(w, apperrors.Unauthorized("authentication required", err))
		return
	}

	req, err := DecodeCreateTaskRequest(r)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid request body", err))
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, apperrors.ValidationError(err.Error(), err))
		return
	}

	result, err := h.usecase.Create(r.Context(), userID, req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, result, nil)
}
