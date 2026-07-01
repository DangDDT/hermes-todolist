package task_update

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/infra/middleware"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Usecase handles task updates.
type Usecase struct {
	taskRepo task.TaskRepository
}

// NewUsecase creates new task update usecase.
func NewUsecase(taskRepo task.TaskRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo}
}

// UpdateTaskResponse response for updated task.
type UpdateTaskResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Priority    string  `json:"priority"`
	DueDate     string  `json:"due_date,omitempty"`
	AssigneeID  *string `json:"assignee_id,omitempty"`
}

// Update modifies existing task (with ownership check).
func (uc *Usecase) Update(ctx context.Context, userID, id uuid.UUID, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	t, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, apperrors.NotFound("Task", err)
	}

	if t.IsDeleted() {
		return nil, apperrors.NotFound("Task", nil)
	}

	// Ownership check: only the creator can update their task
	if t.CreatorID != userID {
		return nil, apperrors.NotFound("Task", nil)
	}

	if req.Title != nil {
		if *req.Title == "" {
			return nil, apperrors.ValidationError("title cannot be empty", nil)
		}
		t.Title = *req.Title
	}
	if req.Description != nil {
		t.Description = *req.Description
	}
	if req.Status != nil {
		newStatus := task.TaskStatus(*req.Status)
		if err := t.UpdateStatus(newStatus); err != nil {
			return nil, apperrors.ValidationError(err.Error(), err)
		}
	}
	if req.Priority != nil {
		newPriority := task.TaskPriority(*req.Priority)
		if err := t.UpdatePriority(newPriority); err != nil {
			return nil, apperrors.ValidationError(err.Error(), err)
		}
	}
	if req.DueDate != nil {
		if *req.DueDate == "" {
			t.DueDate = nil
		} else {
			parsed, err := time.Parse(time.RFC3339, *req.DueDate)
			if err != nil {
				return nil, apperrors.ValidationError("invalid due_date format", err)
			}
			t.DueDate = &parsed
		}
	}
	if req.AssigneeID != nil {
		if *req.AssigneeID == "" {
			t.AssigneeID = nil
		} else {
			assigneeID, err := uuid.Parse(*req.AssigneeID)
			if err != nil {
				return nil, apperrors.ValidationError("invalid assignee_id", err)
			}
			t.AssignTo(assigneeID)
		}
	}

	if err := uc.taskRepo.Update(ctx, t); err != nil {
		return nil, apperrors.Internal(err)
	}

	resp := &UpdateTaskResponse{
		ID:       t.ID.String(),
		Title:    t.Title,
		Description: t.Description,
		Status:   string(t.Status),
		Priority: string(t.Priority),
	}
	if t.DueDate != nil {
		resp.DueDate = t.DueDate.Format(time.RFC3339)
	}
	if t.AssigneeID != nil {
		s := t.AssigneeID.String()
		resp.AssigneeID = &s
	}
	return resp, nil
}

// Handler handles PUT /tasks/{id}.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates new task update handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Update godoc
// @Summary Update task
// @Description Partially update task (only send fields to change) (requires authentication)
// @Tags tasks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Task ID"
// @Param request body UpdateTaskRequest true "Fields to update"
// @Success 200 {object} response.Envelope
// @Failure 400 {object} response.Envelope
// @Failure 401 {object} response.Envelope
// @Failure 404 {object} response.Envelope
// @Router /tasks/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		response.Error(w, apperrors.Unauthorized("invalid token", err))
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid task id", err))
		return
	}

	req, err := DecodeUpdateTaskRequest(r)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid request body", err))
		return
	}

	result, err := h.usecase.Update(r.Context(), userID, id, req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, result, nil)
}
