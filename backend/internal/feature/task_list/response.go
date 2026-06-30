package task_list

import "time"

// TaskItem represents a task in a list.
type TaskItem struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatorID   string     `json:"creator_id"`
	AssigneeID  *string    `json:"assignee_id,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ListTasksResponse is the response for the task list endpoint.
type ListTasksResponse struct {
	Tasks []TaskItem `json:"tasks"`
}
