package task_get

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Usecase handles retrieving a single task.
type Usecase struct {
	taskRepo task.TaskRepository
}

// NewUsecase creates a new task get usecase.
func NewUsecase(taskRepo task.TaskRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo}
}

// GetTaskResponse is the response for a single task.
type GetTaskResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Priority    string  `json:"priority"`
	DueDate     string  `json:"due_date,omitempty"`
	CreatorID   string  `json:"creator_id"`
	AssigneeID  *string `json:"assignee_id,omitempty"`
}

// Get retrieves a task by ID.
func (uc *Usecase) Get(ctx context.Context, id uuid.UUID) (*GetTaskResponse, error) {
	t, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, apperrors.NotFound("Task", err)
	}

	if t.IsDeleted() {
		return nil, apperrors.NotFound("Task", nil)
	}

	resp := &GetTaskResponse{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		Priority:    string(t.Priority),
		CreatorID:   t.CreatorID.String(),
	}
	if t.DueDate != nil {
		resp.DueDate = t.DueDate.Format("2006-01-02T15:04:05Z")
	}
	if t.AssigneeID != nil {
		s := t.AssigneeID.String()
		resp.AssigneeID = &s
	}
	return resp, nil
}

// Handler handles GET /tasks/{id}.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new task get handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Get godoc
// @Summary      Get a task
// @Description  Get a single task by ID (requires authentication)
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Failure      404  {object}  response.Envelope
// @Router       /tasks/{id} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid task id", err))
		return
	}

	result, err := h.usecase.Get(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, result, nil)
}
