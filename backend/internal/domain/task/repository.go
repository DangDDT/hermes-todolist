package task

import (
	"context"

	"github.com/google/uuid"
)

// TaskFilter holds optional filters for listing tasks.
type TaskFilter struct {
	Status     *TaskStatus
	Priority   *TaskPriority
	CreatorID  *uuid.UUID
	AssigneeID *uuid.UUID
	Search     *string
}

// TaskRepository defines the persistence contract for tasks.
type TaskRepository interface {
	Create(ctx context.Context, t *Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*Task, error)
	List(ctx context.Context, filter TaskFilter, offset, limit int) ([]*Task, int, error)
	Update(ctx context.Context, t *Task) error
	SoftDelete(ctx context.Context, id uuid.UUID) error
}
