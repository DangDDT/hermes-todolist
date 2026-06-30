package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

// ErrNotImplemented is returned when a repository method has not been implemented yet.
var ErrNotImplemented = errors.New("repository method not implemented")

// TaskRepo implements task.TaskRepository using PostgreSQL.
type TaskRepo struct {
	pool *pgxpool.Pool
}

// NewTaskRepo creates a new TaskRepo.
func NewTaskRepo(pool *pgxpool.Pool) *TaskRepo {
	return &TaskRepo{pool: pool}
}

// Create persists a new task.
func (r *TaskRepo) Create(ctx context.Context, t *task.Task) error {
	return ErrNotImplemented
}

// GetByID retrieves a task by its ID.
func (r *TaskRepo) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	return nil, ErrNotImplemented
}

// List retrieves tasks with filtering and pagination.
func (r *TaskRepo) List(ctx context.Context, filter task.TaskFilter, offset, limit int) ([]*task.Task, int, error) {
	return nil, 0, ErrNotImplemented
}

// Update persists changes to an existing task.
func (r *TaskRepo) Update(ctx context.Context, t *task.Task) error {
	return ErrNotImplemented
}

// SoftDelete marks a task as deleted.
func (r *TaskRepo) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return ErrNotImplemented
}
