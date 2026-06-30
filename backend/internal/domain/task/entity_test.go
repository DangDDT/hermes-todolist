package task

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewTaskDefaultsToTodoAndPreservesFields(t *testing.T) {
	creatorID := uuid.New()
	dueDate := time.Now().UTC().Add(24 * time.Hour)

	got, err := NewTask("Ship backend", "Finish APIs", PriorityHigh, &dueDate, creatorID)
	if err != nil {
		t.Fatalf("NewTask failed: %v", err)
	}

	if got.Title != "Ship backend" || got.Description != "Finish APIs" {
		t.Fatalf("unexpected task fields: %+v", got)
	}
	if got.Status != StatusTODO {
		t.Fatalf("expected TODO status, got %s", got.Status)
	}
	if got.Priority != PriorityHigh {
		t.Fatalf("expected HIGH priority, got %s", got.Priority)
	}
	if got.CreatorID != creatorID {
		t.Fatalf("expected creator %s, got %s", creatorID, got.CreatorID)
	}
	if got.DueDate == nil || !got.DueDate.Equal(dueDate) {
		t.Fatalf("expected due date %v, got %v", dueDate, got.DueDate)
	}
}

func TestNewTaskValidatesRequiredTitleAndPriority(t *testing.T) {
	if _, err := NewTask("", "", PriorityMedium, nil, uuid.New()); err == nil {
		t.Fatal("expected title validation error")
	}
	if _, err := NewTask("Bad priority", "", TaskPriority("INVALID"), nil, uuid.New()); err == nil {
		t.Fatal("expected priority validation error")
	}
}

func TestTaskStatusTransitions(t *testing.T) {
	task, err := NewTask("Backend", "", PriorityMedium, nil, uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	if err := task.UpdateStatus(StatusInProgress); err != nil {
		t.Fatalf("expected TODO -> IN_PROGRESS to be valid: %v", err)
	}
	if err := task.UpdateStatus(StatusTODO); err == nil {
		t.Fatal("expected IN_PROGRESS -> TODO to be invalid")
	}
	if err := task.UpdateStatus(StatusDone); err != nil {
		t.Fatalf("expected IN_PROGRESS -> DONE to be valid: %v", err)
	}
	if err := task.UpdateStatus(StatusCancelled); err == nil {
		t.Fatal("expected DONE -> CANCELLED to be invalid")
	}
}

func TestTaskAssignmentPriorityAndSoftDelete(t *testing.T) {
	task, err := NewTask("Backend", "", PriorityLow, nil, uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	assignee := uuid.New()

	task.AssignTo(assignee)
	if task.AssigneeID == nil || *task.AssigneeID != assignee {
		t.Fatalf("expected assignee %s, got %v", assignee, task.AssigneeID)
	}
	if err := task.UpdatePriority(PriorityUrgent); err != nil {
		t.Fatalf("expected priority update to pass: %v", err)
	}
	if task.Priority != PriorityUrgent {
		t.Fatalf("expected URGENT priority, got %s", task.Priority)
	}
	task.SoftDelete()
	if !task.IsDeleted() || task.DeletedAt == nil {
		t.Fatal("expected task to be soft deleted")
	}
}
