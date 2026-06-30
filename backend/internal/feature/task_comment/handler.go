package task_comment

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/infra/middleware"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Handler handles task comment requests.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new task comment handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) taskID(r *http.Request) (uuid.UUID, error) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, apperrors.ValidationError("invalid task id", err)
	}
	return id, nil
}

// List handles GET /tasks/{id}/comments.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	taskID, err := h.taskID(r)
	if err != nil {
		response.Error(w, err)
		return
	}

	result, err := h.usecase.List(r.Context(), taskID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, result, nil)
}

// Create handles POST /tasks/{id}/comments.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	taskID, err := h.taskID(r)
	if err != nil {
		response.Error(w, err)
		return
	}

	authorID, err := middleware.GetUserID(r.Context())
	if err != nil {
		response.Error(w, apperrors.Unauthorized("authentication required", err))
		return
	}

	req, err := DecodeCreateCommentRequest(r)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid request body", err))
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, apperrors.ValidationError(err.Error(), err))
		return
	}

	result, err := h.usecase.Create(r.Context(), taskID, authorID, req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusCreated, result, nil)
}
