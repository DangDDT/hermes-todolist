package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Task is the aggregate root for tasks.
type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      TaskStatus
	Priority    TaskPriority
	DueDate     *time.Time
	CreatorID   uuid.UUID
	AssigneeID  *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// NewTask creates a new Task with validation.
func NewTask(title, description string, priority TaskPriority, dueDate *time.Time, creatorID uuid.UUID) (*Task, error) {
	if title == "" {
		return nil, errors.New("task title is required")
	}
	if !priority.IsValid() {
		return nil, errors.New("invalid task priority")
	}
	now := time.Now().UTC()
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      StatusTODO,
		Priority:    priority,
		DueDate:     dueDate,
		CreatorID:   creatorID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// AssignTo assigns the task to a user.
func (t *Task) AssignTo(userID uuid.UUID) {
	t.AssigneeID = &userID
	t.UpdatedAt = time.Now().UTC()
}

// UpdateStatus updates the task status with transition validation.
func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if !newStatus.IsValid() {
		return errors.New("invalid task status")
	}
	if !IsValidTransition(t.Status, newStatus) {
		return errors.New("invalid status transition")
	}
	t.Status = newStatus
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// UpdatePriority updates the task priority.
func (t *Task) UpdatePriority(priority TaskPriority) error {
	if !priority.IsValid() {
		return errors.New("invalid task priority")
	}
	t.Priority = priority
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// SoftDelete marks the task as deleted.
func (t *Task) SoftDelete() {
	now := time.Now().UTC()
	t.DeletedAt = &now
	t.UpdatedAt = now
}

// IsDeleted returns true if the task has been soft-deleted.
func (t *Task) IsDeleted() bool {
	return t.DeletedAt != nil
}
