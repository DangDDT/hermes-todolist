package task_update

import (
	"encoding/json"
	"net/http"
)

// UpdateTaskRequest is the request body for updating a task.
type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	Priority    *string `json:"priority"`
	DueDate     *string `json:"due_date"`
	AssigneeID  *string `json:"assignee_id"`
}

// DecodeUpdateTaskRequest parses the JSON request body.
func DecodeUpdateTaskRequest(r *http.Request) (*UpdateTaskRequest, error) {
	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
