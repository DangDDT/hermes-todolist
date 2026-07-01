package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

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
	_, err := r.pool.Exec(ctx,
		`INSERT INTO tasks (id, title, description, status, priority, due_date, creator_id, assignee_id, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		t.ID, t.Title, t.Description, string(t.Status), string(t.Priority),
		t.DueDate, t.CreatorID, t.AssigneeID, t.CreatedAt, t.UpdatedAt,
	)
	return err
}

// GetByID retrieves a single task by ID (not soft-deleted).
func (r *TaskRepo) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	t := &task.Task{}
	var dueDate *time.Time
	err := r.pool.QueryRow(ctx,
		`SELECT id, title, description, status, priority, due_date, creator_id, assignee_id, created_at, updated_at
		 FROM tasks WHERE id = $1 AND deleted_at IS NULL`, id,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority,
		&dueDate, &t.CreatorID, &t.AssigneeID, &t.CreatedAt, &t.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	t.DueDate = dueDate
	return t, nil
}

// List queries tasks with filters and pagination.
func (r *TaskRepo) List(ctx context.Context, filter task.TaskFilter, offset, limit int) ([]*task.Task, int, error) {
	// Count total matching
	var total int
	countQuery := `SELECT COUNT(*) FROM tasks WHERE deleted_at IS NULL`
	countArgs := []interface{}{}
	argIdx := 1

	if filter.Status != nil {
		countQuery += ` AND status = $` + itoa(argIdx)
		countArgs = append(countArgs, string(*filter.Status))
		argIdx++
	}
	if filter.AssigneeID != nil {
		countQuery += ` AND assignee_id = $` + itoa(argIdx)
		countArgs = append(countArgs, *filter.AssigneeID)
		argIdx++
	}
	if filter.CreatorID != nil {
		countQuery += ` AND creator_id = $` + itoa(argIdx)
		countArgs = append(countArgs, *filter.CreatorID)
		argIdx++
	}
	if filter.Search != nil && *filter.Search != "" {
		countQuery += ` AND (title ILIKE $` + itoa(argIdx) + ` OR description ILIKE $` + itoa(argIdx) + `)`
		searchPattern := "%" + *filter.Search + "%"
		countArgs = append(countArgs, searchPattern)
		argIdx++
	}

	err := r.pool.QueryRow(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Build data query
	argIdx = 1
	query := `SELECT id, title, description, status, priority, due_date, creator_id, assignee_id, created_at, updated_at
		FROM tasks WHERE deleted_at IS NULL`
	args := []interface{}{}

	if filter.Status != nil {
		query += ` AND status = $` + itoa(argIdx)
		args = append(args, string(*filter.Status))
		argIdx++
	}
	if filter.AssigneeID != nil {
		query += ` AND assignee_id = $` + itoa(argIdx)
		args = append(args, *filter.AssigneeID)
		argIdx++
	}
	if filter.CreatorID != nil {
		query += ` AND creator_id = $` + itoa(argIdx)
		args = append(args, *filter.CreatorID)
		argIdx++
	}
	if filter.Search != nil && *filter.Search != "" {
		query += ` AND (title ILIKE $` + itoa(argIdx) + ` OR description ILIKE $` + itoa(argIdx) + `)`
		searchPattern := "%" + *filter.Search + "%"
		args = append(args, searchPattern)
		argIdx++
	}

	query += ` ORDER BY created_at DESC`
	query += ` LIMIT $` + itoa(argIdx) + ` OFFSET $` + itoa(argIdx+1)
	args = append(args, limit, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		t := &task.Task{}
		var dueDate *time.Time
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority,
			&dueDate, &t.CreatorID, &t.AssigneeID, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, 0, err
		}
		t.DueDate = dueDate
		tasks = append(tasks, t)
	}

	return tasks, total, nil
}

// Update modifies an existing task.
func (r *TaskRepo) Update(ctx context.Context, t *task.Task) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE tasks SET title=$1, description=$2, status=$3, priority=$4, due_date=$5, assignee_id=$6, updated_at=$7
		 WHERE id=$8 AND deleted_at IS NULL`,
		t.Title, t.Description, string(t.Status), string(t.Priority),
		t.DueDate, t.AssigneeID, t.UpdatedAt, t.ID,
	)
	return err
}

// SoftDelete marks a task as deleted.
func (r *TaskRepo) SoftDelete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE tasks SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`, id,
	)
	return err
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
