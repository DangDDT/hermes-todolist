package comment

import (
	"context"

	"github.com/google/uuid"
)

// CommentRepository defines the persistence contract for task comments.
type CommentRepository interface {
	Create(ctx context.Context, c *Comment) error
	GetByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	ListByTask(ctx context.Context, taskID uuid.UUID) ([]*Comment, error)
}
