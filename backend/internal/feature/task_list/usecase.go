package task_list

import (
	"context"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/pagination"
)

// Usecase handles listing tasks.
type Usecase struct {
	taskRepo task.TaskRepository
}

// NewUsecase creates a new task list usecase.
func NewUsecase(taskRepo task.TaskRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo}
}

// List retrieves tasks with filters and pagination.
func (uc *Usecase) List(ctx context.Context, creatorID uuid.UUID, filter *ListTasksRequest, page, perPage, offset int) (*ListTasksResponse, pagination.Meta, error) {
	taskFilter := task.TaskFilter{
		CreatorID: &creatorID,
	}

	if filter.Status != "" {
		s := task.TaskStatus(filter.Status)
		if !s.IsValid() {
			return nil, pagination.Meta{}, apperrors.ValidationError("invalid status filter", nil)
		}
		taskFilter.Status = &s
	}
	if filter.Priority != "" {
		p := task.TaskPriority(filter.Priority)
		if !p.IsValid() {
			return nil, pagination.Meta{}, apperrors.ValidationError("invalid priority filter", nil)
		}
		taskFilter.Priority = &p
	}
	if filter.Search != "" {
		taskFilter.Search = &filter.Search
	}

	tasks, total, err := uc.taskRepo.List(ctx, taskFilter, offset, perPage)
	if err != nil {
		return nil, pagination.Meta{}, apperrors.Internal(err)
	}

	items := make([]TaskItem, 0, len(tasks))
	for _, t := range tasks {
		item := TaskItem{
			ID:          t.ID.String(),
			Title:       t.Title,
			Description: t.Description,
			Status:      string(t.Status),
			Priority:    string(t.Priority),
			DueDate:     t.DueDate,
			CreatorID:   t.CreatorID.String(),
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}
		if t.AssigneeID != nil {
			s := t.AssigneeID.String()
			item.AssigneeID = &s
		}
		items = append(items, item)
	}

	meta := pagination.NewMeta(page, perPage, total)

	return &ListTasksResponse{Tasks: items}, meta, nil
}
