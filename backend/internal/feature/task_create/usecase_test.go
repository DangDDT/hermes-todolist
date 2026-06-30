package task_create

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct { create func(context.Context, *task.Task) error }
func (s taskRepoStub) Create(ctx context.Context, t *task.Task) error { if s.create != nil { return s.create(ctx,t) }; return nil }
func (s taskRepoStub) GetByID(context.Context, uuid.UUID) (*task.Task,error) { return nil,nil }
func (s taskRepoStub) List(context.Context, task.TaskFilter, int, int) ([]*task.Task,int,error) { return nil,0,nil }
func (s taskRepoStub) Update(context.Context, *task.Task) error { return nil }
func (s taskRepoStub) SoftDelete(context.Context, uuid.UUID) error { return nil }

func TestCreatePersistsTaskWithDefaultsAndAssignee(t *testing.T) {
	creatorID := uuid.New()
	assigneeID := uuid.New()
	dueDate := time.Now().UTC().Truncate(time.Second).Add(48*time.Hour)
	var saved *task.Task
	uc := NewUsecase(taskRepoStub{create: func(ctx context.Context, t *task.Task) error { saved=t; return nil }})

	resp, err := uc.Create(context.Background(), creatorID, &CreateTaskRequest{Title:"Backend", Description:"Finish", Priority:"HIGH", DueDate: dueDate.Format(time.RFC3339), AssigneeID: assigneeID.String()})
	if err != nil { t.Fatalf("create failed: %v", err) }
	if saved == nil { t.Fatal("expected task to be persisted") }
	if saved.CreatorID != creatorID || saved.Status != task.StatusTODO || saved.Priority != task.PriorityHigh { t.Fatalf("unexpected saved task: %+v", saved) }
	if saved.AssigneeID == nil || *saved.AssigneeID != assigneeID { t.Fatalf("unexpected assignee: %v", saved.AssigneeID) }
	if resp.Title != "Backend" || resp.AssigneeID == nil || *resp.AssigneeID != assigneeID.String() { t.Fatalf("unexpected response: %+v", resp) }
}

func TestCreateRejectsInvalidPriorityAndAssignee(t *testing.T) {
	uc := NewUsecase(taskRepoStub{create: func(ctx context.Context, created *task.Task) error { t.Fatal("should not persist invalid task"); return nil }})
	_, err := uc.Create(context.Background(), uuid.New(), &CreateTaskRequest{Title:"Backend", Priority:"BAD"})
	if err == nil || !strings.Contains(err.Error(), "invalid priority") { t.Fatalf("expected invalid priority, got %v", err) }
	_, err = uc.Create(context.Background(), uuid.New(), &CreateTaskRequest{Title:"Backend", Priority:"HIGH", AssigneeID:"not-a-uuid"})
	if err == nil || !strings.Contains(err.Error(), "invalid assignee_id") { t.Fatalf("expected invalid assignee, got %v", err) }
}
