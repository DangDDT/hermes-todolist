package task_delete

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct {
	getByID     func(context.Context, uuid.UUID) (*task.Task, error)
	softDelete  func(context.Context, uuid.UUID) error
}

func (s taskRepoStub) Create(context.Context, *task.Task) error                          { return nil }
func (s taskRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	if s.getByID != nil {
		return s.getByID(ctx, id)
	}
	return nil, nil
}
func (s taskRepoStub) List(context.Context, task.TaskFilter, int, int) ([]*task.Task, int, error) { return nil, 0, nil }
func (s taskRepoStub) Update(context.Context, *task.Task) error       { return nil }
func (s taskRepoStub) SoftDelete(ctx context.Context, id uuid.UUID) error {
	if s.softDelete != nil {
		return s.softDelete(ctx, id)
	}
	return nil
}

func TestDeleteSoftDeletesExistingTask(t *testing.T) {
	creatorID := uuid.New()
	stored, err := task.NewTask("Backend", "", task.PriorityMedium, nil, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	called := false
	uc := NewUsecase(taskRepoStub{
		getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil },
		softDelete: func(ctx context.Context, id uuid.UUID) error {
			called = true
			if id != stored.ID {
				t.Fatalf("unexpected id: %s", id)
			}
			return nil
		},
	})

	if err := uc.Delete(context.Background(), creatorID, stored.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if !called {
		t.Fatal("expected SoftDelete to be called")
	}
}

func TestDeleteRejectsAlreadyDeletedTask(t *testing.T) {
	creatorID := uuid.New()
	stored, err := task.NewTask("Backend", "", task.PriorityMedium, nil, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	stored.SoftDelete()
	uc := NewUsecase(taskRepoStub{
		getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil },
		softDelete: func(ctx context.Context, id uuid.UUID) error {
			t.Fatal("should not soft delete twice")
			return nil
		},
	})

	err = uc.Delete(context.Background(), creatorID, stored.ID)
	if err == nil || !strings.Contains(err.Error(), "NOT_FOUND") {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestDeleteRejectsOtherUsersTask(t *testing.T) {
	creatorID := uuid.New()
	otherUserID := uuid.New()
	stored, err := task.NewTask("My Task", "", task.PriorityMedium, nil, creatorID)
	if err != nil {
		t.Fatal(err)
	}
	uc := NewUsecase(taskRepoStub{
		getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) { return stored, nil },
	})

	err = uc.Delete(context.Background(), otherUserID, stored.ID)
	if err == nil || !strings.Contains(err.Error(), "NOT_FOUND") {
		t.Fatalf("expected not found for other user's task, got %v", err)
	}
}
