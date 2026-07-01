package task_create

import (
	"context"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
)

// Usecase handles task creation.
type Usecase struct {
	taskRepo task.TaskRepository
}

// NewUsecase creates new task create usecase.
func NewUsecase(taskRepo task.TaskRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo}
}

// Create creates new task.
func (uc *Usecase) Create(ctx context.Context, creatorID uuid.UUID, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	priority := task.PriorityMedium
	if req.Priority != "" {
		p := task.TaskPriority(req.Priority)
		if !p.IsValid() {
			return nil, apperrors.ValidationError("invalid priority", nil)
		}
		priority = p
	}

	status := task.StatusTODO
	if req.Status != "" {
		s := task.TaskStatus(req.Status)
		if !s.IsValid() {
			return nil, apperrors.ValidationError("invalid status", nil)
		}
		status = s
	}

	dueDate := req.ParseDueDate()

	t, err := task.NewTask(req.Title, req.Description, priority, dueDate, creatorID)
	if err != nil {
		return nil, apperrors.ValidationError(err.Error(), err)
	}
	// Apply status from request (NewTask defaults to StatusTodo)
	t.Status = status

	// Assign assignee if provided.
	if req.AssigneeID != "" {
		assigneeID, err := uuid.Parse(req.AssigneeID)
		if err != nil {
			return nil, apperrors.ValidationError("invalid assignee_id", err)
		}
		t.AssignTo(assigneeID)
	}

	if err := uc.taskRepo.Create(ctx, t); err != nil {
		return nil, apperrors.Internal(err)
	}

	resp := &CreateTaskResponse{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		Priority:    string(t.Priority),
		DueDate:     t.DueDate,
		CreatorID:   t.CreatorID.String(),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
	if t.AssigneeID != nil {
		s := t.AssigneeID.String()
		resp.AssigneeID = &s
	}
	return resp, nil
}
