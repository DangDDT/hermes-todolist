package task_comment

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/DangDDT/hermes-todolist/backend/internal/domain/comment"
	"github.com/DangDDT/hermes-todolist/backend/internal/domain/task"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
)

// Usecase handles task comment operations.
type Usecase struct {
	taskRepo    task.TaskRepository
	commentRepo comment.CommentRepository
}

// NewUsecase creates a new task comment usecase.
func NewUsecase(taskRepo task.TaskRepository, commentRepo comment.CommentRepository) *Usecase {
	return &Usecase{taskRepo: taskRepo, commentRepo: commentRepo}
}

// Create creates a new comment on a task.
func (uc *Usecase) Create(ctx context.Context, taskID, authorID uuid.UUID, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	if req == nil {
		return nil, apperrors.ValidationError("invalid request body", nil)
	}
	body := strings.TrimSpace(req.Body)
	if body == "" {
		return nil, apperrors.ValidationError("body is required", nil)
	}
	if len(body) > 2000 {
		return nil, apperrors.ValidationError("body must be 2000 characters or less", nil)
	}

	if _, err := uc.taskRepo.GetByID(ctx, taskID); err != nil {
		return nil, apperrors.NotFound("Task", err)
	}

	c := &comment.Comment{
		ID:        uuid.New(),
		TaskID:    taskID,
		AuthorID:  authorID,
		Body:      body,
		CreatedAt: time.Now().UTC(),
	}

	if err := uc.commentRepo.Create(ctx, c); err != nil {
		return nil, apperrors.Internal(err)
	}

	created, err := uc.commentRepo.GetByID(ctx, c.ID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	return &CreateCommentResponse{Comment: toCommentResponse(created)}, nil
}

// List lists comments for a task.
func (uc *Usecase) List(ctx context.Context, taskID uuid.UUID) (*ListCommentsResponse, error) {
	if _, err := uc.taskRepo.GetByID(ctx, taskID); err != nil {
		return nil, apperrors.NotFound("Task", err)
	}

	comments, err := uc.commentRepo.ListByTask(ctx, taskID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	resp := &ListCommentsResponse{Comments: make([]CommentResponse, 0, len(comments))}
	for _, c := range comments {
		resp.Comments = append(resp.Comments, toCommentResponse(c))
	}
	return resp, nil
}
