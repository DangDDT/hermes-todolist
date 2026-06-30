package task_comment

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/comment"
	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct {
	getByID func(ctx context.Context, id uuid.UUID) (*task.Task, error)
}

func (s taskRepoStub) Create(ctx context.Context, t *task.Task) error { return nil }
func (s taskRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	return s.getByID(ctx, id)
}
func (s taskRepoStub) List(ctx context.Context, filter task.TaskFilter, offset, limit int) ([]*task.Task, int, error) {
	return nil, 0, nil
}
func (s taskRepoStub) Update(ctx context.Context, t *task.Task) error { return nil }
func (s taskRepoStub) SoftDelete(ctx context.Context, id uuid.UUID) error { return nil }

type commentRepoStub struct {
	create  func(ctx context.Context, c *comment.Comment) error
	getByID func(ctx context.Context, id uuid.UUID) (*comment.Comment, error)
	list    func(ctx context.Context, taskID uuid.UUID) ([]*comment.Comment, error)
}

func (s commentRepoStub) Create(ctx context.Context, c *comment.Comment) error { return s.create(ctx, c) }
func (s commentRepoStub) GetByID(ctx context.Context, id uuid.UUID) (*comment.Comment, error) {
	return s.getByID(ctx, id)
}
func (s commentRepoStub) ListByTask(ctx context.Context, taskID uuid.UUID) ([]*comment.Comment, error) {
	return s.list(ctx, taskID)
}

func TestCreateRejectsEmptyBody(t *testing.T) {
	uc := NewUsecase(taskRepoStub{}, commentRepoStub{})
	_, err := uc.Create(context.Background(), uuid.New(), uuid.New(), &CreateCommentRequest{Body: ""})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCreatePersistsCommentAfterTaskLookup(t *testing.T) {
	taskID := uuid.New()
	authorID := uuid.New()
	wantBody := "Looks good to me"
	seen := false

	uc := NewUsecase(
		taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) {
			if id != taskID {
				t.Fatalf("unexpected task id: %s", id)
			}
			return &task.Task{ID: taskID}, nil
		}},
		commentRepoStub{create: func(ctx context.Context, c *comment.Comment) error {
			seen = true
			if c.TaskID != taskID {
				t.Fatalf("unexpected task id: %s", c.TaskID)
			}
			if c.AuthorID != authorID {
				t.Fatalf("unexpected author id: %s", c.AuthorID)
			}
			if c.Body != wantBody {
				t.Fatalf("unexpected body: %q", c.Body)
			}
			if c.ID == uuid.Nil {
				t.Fatal("expected generated id")
			}
			if c.CreatedAt.IsZero() {
				t.Fatal("expected created at")
			}
			return nil
		}, getByID: func(ctx context.Context, id uuid.UUID) (*comment.Comment, error) {
			if !seen {
				t.Fatal("expected create before get")
			}
			return &comment.Comment{ID: id, TaskID: taskID, AuthorID: authorID, AuthorName: "DangDDT", Body: wantBody, CreatedAt: time.Now().UTC()}, nil
		}},
	)

	resp, err := uc.Create(context.Background(), taskID, authorID, &CreateCommentRequest{Body: wantBody})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !seen {
		t.Fatal("expected repository create to be called")
	}
	if resp.Comment.Body != wantBody {
		t.Fatalf("unexpected response body: %q", resp.Comment.Body)
	}
}

func TestListReturnsCommentsForTask(t *testing.T) {
	taskID := uuid.New()
	now := time.Now().UTC()
	uc := NewUsecase(
		taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) {
			if id != taskID {
				t.Fatalf("unexpected task id: %s", id)
			}
			return &task.Task{ID: taskID}, nil
		}},
		commentRepoStub{list: func(ctx context.Context, id uuid.UUID) ([]*comment.Comment, error) {
			if id != taskID {
				t.Fatalf("unexpected task id: %s", id)
			}
			return []*comment.Comment{{ID: uuid.New(), TaskID: taskID, AuthorID: uuid.New(), AuthorName: "DangDDT", Body: "First", CreatedAt: now}}, nil
		}},
	)

	resp, err := uc.List(context.Background(), taskID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Comments) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(resp.Comments))
	}
	if resp.Comments[0].AuthorName != "DangDDT" {
		t.Fatalf("unexpected author name: %q", resp.Comments[0].AuthorName)
	}
}

func TestListReturnsNotFoundWhenTaskMissing(t *testing.T) {
	uc := NewUsecase(
		taskRepoStub{getByID: func(ctx context.Context, id uuid.UUID) (*task.Task, error) {
			return nil, errors.New("boom")
		}},
		commentRepoStub{list: func(ctx context.Context, taskID uuid.UUID) ([]*comment.Comment, error) {
			return nil, nil
		}},
	)

	_, err := uc.List(context.Background(), uuid.New())
	if err == nil {
		t.Fatal("expected not found error")
	}
}
