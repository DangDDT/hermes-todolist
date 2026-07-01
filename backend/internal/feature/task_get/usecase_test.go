package task_get

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct{ getByID func(context.Context, uuid.UUID) (*task.Task, error) }
func (s taskRepoStub) Create(context.Context, *task.Task) error          { return nil }
func (s taskRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	if s.getByID != nil {
		return s.getByID(ctx, id)
	}
	return nil, nil
}
func (s taskRepoStub) List(context.Context, task.TaskFilter, int, int) ([]*task.Task, int, error) { return nil, 0, nil }
func (s taskRepoStub) Update(context.Context, *task.Task) error       { return nil }
func (s taskRepoStub) SoftDelete(context.Context, uuid.UUID) error    { return nil }

func TestGetReturnsTaskResponse(t *testing.T) {
	creatorID := uuid.New()
	assigneeID := uuid.New()
	dueDate := time.Now().UTC().Truncate(time.Second)
	stored, err := task.NewTask("Backend", "Finish", task.PriorityHigh, &dueDate, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	stored.AssignTo(assigneeID)
	uc := NewUsecase(taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil }})

	resp, err := uc.Get(context.Background(), creatorID, stored.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if resp.ID != stored.ID.String() || resp.Title != "Backend" || resp.Priority != "HIGH" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if resp.AssigneeID == nil || *resp.AssigneeID != assigneeID.String() {
		t.Fatalf("unexpected assignee: %+v", resp.AssigneeID)
	}
	if resp.DueDate == "" {
		t.Fatal("expected due date formatted")
	}
}

func TestGetHidesSoftDeletedTask(t *testing.T) {
	creatorID := uuid.New()
	stored, err := task.NewTask("Backend", "", task.PriorityMedium, nil, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	stored.SoftDelete()
	uc := NewUsecase(taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil }})

	_, err = uc.Get(context.Background(), creatorID, stored.ID)
	if err == nil || !strings.Contains(err.Error(), "NOT_FOUND") {
		t.Fatalf("expected not found for deleted task, got %v", err)
	}
}

func TestGetRejectsOtherUsersTask(t *testing.T) {
	creatorID := uuid.New()
	otherUserID := uuid.New()
	stored, err := task.NewTask("My Task", "", task.PriorityLow, nil, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	uc := NewUsecase(taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil }})

	_, err = uc.Get(context.Background(), otherUserID, stored.ID)
	if err == nil || !strings.Contains(err.Error(), "NOT_FOUND") {
		t.Fatalf("expected not found for other user's task, got %v", err)
	}
}
