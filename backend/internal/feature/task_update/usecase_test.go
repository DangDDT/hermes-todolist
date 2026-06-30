package task_update

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct {
	getByID func(context.Context, uuid.UUID) (*task.Task, error)
	update  func(context.Context, *task.Task) error
}
func (s taskRepoStub) Create(context.Context, *task.Task) error { return nil }
func (s taskRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*task.Task,error) { if s.getByID != nil { return s.getByID(ctx,id) }; return nil,nil }
func (s taskRepoStub) List(context.Context, task.TaskFilter, int, int) ([]*task.Task,int,error) { return nil,0,nil }
func (s taskRepoStub) Update(ctx context.Context, task *task.Task) error { if s.update != nil { return s.update(ctx,task) }; return nil }
func (s taskRepoStub) SoftDelete(context.Context, uuid.UUID) error { return nil }

func ptr(s string) *string { return &s }

func TestUpdateModifiesAllowedFieldsAndPersists(t *testing.T) {
	stored, err := task.NewTask("Backend", "Old", task.PriorityMedium, nil, uuid.New())
	if err != nil { t.Fatal(err) }
	var saved *task.Task
	uc := NewUsecase(taskRepoStub{
		getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil },
		update: func(ctx context.Context, updated *task.Task) error { saved = updated; return nil },
	})
	due := time.Now().UTC().Truncate(time.Second).Add(24*time.Hour).Format(time.RFC3339)
	resp, err := uc.Update(context.Background(), stored.ID, &UpdateTaskRequest{Title: ptr("Backend done"), Description: ptr("New"), Status: ptr("IN_PROGRESS"), Priority: ptr("HIGH"), DueDate: &due})
	if err != nil { t.Fatalf("update failed: %v", err) }
	if saved == nil { t.Fatal("expected updated task to be persisted") }
	if saved.Title != "Backend done" || saved.Description != "New" || saved.Status != task.StatusInProgress || saved.Priority != task.PriorityHigh { t.Fatalf("unexpected saved task: %+v", saved) }
	if resp.Title != "Backend done" || resp.Status != "IN_PROGRESS" || resp.Priority != "HIGH" { t.Fatalf("unexpected response: %+v", resp) }
}

func TestUpdateRejectsInvalidStatusTransition(t *testing.T) {
	stored, err := task.NewTask("Backend", "", task.PriorityMedium, nil, uuid.New())
	if err != nil { t.Fatal(err) }
	uc := NewUsecase(taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil }})

	_, err = uc.Update(context.Background(), stored.ID, &UpdateTaskRequest{Status: ptr("DONE")})
	if err == nil || !strings.Contains(err.Error(), "invalid status transition") { t.Fatalf("expected invalid transition, got %v", err) }
}
