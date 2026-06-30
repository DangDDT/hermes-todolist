package task_delete

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Usecase handles task deletion.
type Usecase struct {
	taskRepo task.TaskRepository
}

// NewUsecase creates a new task delete usecase.
func NewUsecase(taskRepo task.TaskRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo}
}

// Delete soft-deletes a task.
func (uc *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
	t, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return apperrors.NotFound("Task", err)
	}
	if t.IsDeleted() {
		return apperrors.NotFound("Task", nil)
	}
	if err := uc.taskRepo.SoftDelete(ctx, id); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

// Handler handles DELETE /tasks/{id}.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new task delete handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// Delete handles the delete task request.
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, apperrors.ValidationError("invalid task id", err))
		return
	}

	if err := h.usecase.Delete(r.Context(), id); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusNoContent, nil, nil)
}
