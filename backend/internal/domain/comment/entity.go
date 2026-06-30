package comment

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents a task comment.
type Comment struct {
	ID         uuid.UUID
	TaskID     uuid.UUID
	AuthorID   uuid.UUID
	AuthorName string
	Body       string
	CreatedAt  time.Time
}
