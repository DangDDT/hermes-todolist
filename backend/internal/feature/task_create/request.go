package task_create

import (
	"encoding/json"
	"net/http"
	"time"
)

// CreateTaskRequest request body creating task.
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	AssigneeID  string `json:"assignee_id"`
}

// DecodeCreateTaskRequest parses the JSON request body.
func DecodeCreateTaskRequest(r *http.Request) (*CreateTaskRequest, error) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// Validate checks the request fields.
func (r *CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return &fieldError{"title", "title is required"}
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

// ParseDueDate parses the due date string into *time.Time.
func (r *CreateTaskRequest) ParseDueDate() *time.Time {
	if r.DueDate == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, r.DueDate)
	if err != nil {
		return nil
	}
	return &t
}
