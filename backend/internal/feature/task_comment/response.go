package task_comment

import (
	"time"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/comment"
)

// CommentResponse is the API representation of a comment.
type CommentResponse struct {
	ID         string `json:"id"`
	TaskID     string `json:"task_id"`
	AuthorID   string `json:"author_id"`
	AuthorName string `json:"author_name"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
}

// ListCommentsResponse is the response for listing comments.
type ListCommentsResponse struct {
	Comments []CommentResponse `json:"comments"`
}

// CreateCommentResponse is the response for creating a comment.
type CreateCommentResponse struct {
	Comment CommentResponse `json:"comment"`
}

func toCommentResponse(c *comment.Comment) CommentResponse {
	return CommentResponse{
		ID:         c.ID.String(),
		TaskID:     c.TaskID.String(),
		AuthorID:   c.AuthorID.String(),
		AuthorName: c.AuthorName,
		Body:       c.Body,
		CreatedAt:  c.CreatedAt.UTC().Format(time.RFC3339),
	}
}
