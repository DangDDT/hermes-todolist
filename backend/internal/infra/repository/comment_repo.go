package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/comment"
)

// CommentRepo implements comment.CommentRepository using PostgreSQL.
type CommentRepo struct {
	pool *pgxpool.Pool
}

// NewCommentRepo creates a new CommentRepo.
func NewCommentRepo(pool *pgxpool.Pool) *CommentRepo {
	return &CommentRepo{pool: pool}
}

// Create persists a new comment.
func (r *CommentRepo) Create(ctx context.Context, c *comment.Comment) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO task_comments (id, task_id, author_id, body, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		c.ID, c.TaskID, c.AuthorID, c.Body, c.CreatedAt,
	)
	return err
}

// GetByID retrieves a comment by ID.
func (r *CommentRepo) GetByID(ctx context.Context, id uuid.UUID) (*comment.Comment, error) {
	c := &comment.Comment{}
	var authorName string
	var createdAt time.Time
	if err := r.pool.QueryRow(ctx,
		`SELECT c.id, c.task_id, c.author_id, u.display_name, c.body, c.created_at
		 FROM task_comments c
		 JOIN users u ON u.id = c.author_id
		 WHERE c.id = $1`, id,
	).Scan(&c.ID, &c.TaskID, &c.AuthorID, &authorName, &c.Body, &createdAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	c.AuthorName = authorName
	c.CreatedAt = createdAt
	return c, nil
}

// ListByTask lists comments for a task ordered chronologically.
func (r *CommentRepo) ListByTask(ctx context.Context, taskID uuid.UUID) ([]*comment.Comment, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT c.id, c.task_id, c.author_id, u.display_name, c.body, c.created_at
		 FROM task_comments c
		 JOIN users u ON u.id = c.author_id
		 WHERE c.task_id = $1
		 ORDER BY c.created_at ASC`, taskID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]*comment.Comment, 0)
	for rows.Next() {
		c := &comment.Comment{}
		if err := rows.Scan(&c.ID, &c.TaskID, &c.AuthorID, &c.AuthorName, &c.Body, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}
