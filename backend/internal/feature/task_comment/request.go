package task_comment

import (
	"encoding/json"
	"net/http"
	"strings"
)

// CreateCommentRequest is the request body for creating a task comment.
type CreateCommentRequest struct {
	Body string `json:"body"`
}

// DecodeCreateCommentRequest parses the JSON request body.
func DecodeCreateCommentRequest(r *http.Request) (*CreateCommentRequest, error) {
	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// Validate checks the request body.
func (r *CreateCommentRequest) Validate() error {
	if strings.TrimSpace(r.Body) == "" {
		return &fieldError{Field: "body", Message: "body is required"}
	}
	if len(strings.TrimSpace(r.Body)) > 2000 {
		return &fieldError{Field: "body", Message: "body must be 2000 characters or less"}
	}
	return nil
}

type fieldError struct {
	Field   string
	Message string
}

func (e *fieldError) Error() string {
	return e.Field + ": " + e.Message
}
