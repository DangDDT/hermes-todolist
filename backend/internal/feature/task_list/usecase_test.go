package task_list

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
)

type taskRepoStub struct { list func(context.Context, task.TaskFilter, int, int) ([]*task.Task, int, error) }
func (s taskRepoStub) Create(context.Context, *task.Task) error { return nil }
func (s taskRepoStub) GetByID(context.Context, uuid.UUID) (*task.Task,error) { return nil,nil }
func (s taskRepoStub) List(ctx context.Context, f task.TaskFilter, offset, limit int) ([]*task.Task,int,error) { if s.list != nil { return s.list(ctx,f,offset,limit) }; return nil,0,nil }
func (s taskRepoStub) Update(context.Context, *task.Task) error { return nil }
func (s taskRepoStub) SoftDelete(context.Context, uuid.UUID) error { return nil }

func TestListBuildsFiltersAndPagination(t *testing.T) {
	creatorID := uuid.New()
	stored, err := task.NewTask("Backend", "Finish", task.PriorityHigh, nil, creatorID)
	if err != nil { t.Fatal(err) }
	uc := NewUsecase(taskRepoStub{list: func(ctx context.Context, f task.TaskFilter, offset, limit int) ([]*task.Task, int, error) {
		if f.CreatorID == nil || *f.CreatorID != creatorID { t.Fatalf("unexpected creator filter: %+v", f.CreatorID) }
		if f.Status == nil || *f.Status != task.StatusTODO { t.Fatalf("unexpected status filter: %+v", f.Status) }
		if f.Priority == nil || *f.Priority != task.PriorityHigh { t.Fatalf("unexpected priority filter: %+v", f.Priority) }
		if f.Search == nil || *f.Search != "back" { t.Fatalf("unexpected search filter: %+v", f.Search) }
		if offset != 10 || limit != 5 { t.Fatalf("unexpected pagination offset=%d limit=%d", offset, limit) }
		return []*task.Task{stored}, 12, nil
	}})

	resp, meta, err := uc.List(context.Background(), creatorID, &ListTasksRequest{Status:"TODO", Priority:"HIGH", Search:"back"}, 3, 5, 10)
	if err != nil { t.Fatalf("list failed: %v", err) }
	if len(resp.Tasks) != 1 || resp.Tasks[0].Title != "Backend" { t.Fatalf("unexpected response: %+v", resp) }
	if meta.Page != 3 || meta.PerPage != 5 || meta.Total != 12 { t.Fatalf("unexpected meta: %+v", meta) }
}

func TestListRejectsInvalidFilters(t *testing.T) {
	uc := NewUsecase(taskRepoStub{})
	_, _, err := uc.List(context.Background(), uuid.New(), &ListTasksRequest{Status:"BAD"}, 1, 10, 0)
	if err == nil || !strings.Contains(err.Error(), "invalid status") { t.Fatalf("expected status error, got %v", err) }
	_, _, err = uc.List(context.Background(), uuid.New(), &ListTasksRequest{Priority:"BAD"}, 1, 10, 0)
	if err == nil || !strings.Contains(err.Error(), "invalid priority") { t.Fatalf("expected priority error, got %v", err) }
}
