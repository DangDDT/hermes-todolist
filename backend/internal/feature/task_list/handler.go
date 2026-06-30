package task_list

import (
	"net/http"

	"github.com/DangDDT/hermes-todolist/backend/internal/infra/middleware"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/pagination"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// Handler handles GET /tasks.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new task list handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// List handles the list tasks request.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		response.Error(w, apperrors.Unauthorized("authentication required", err))
		return
	}

	page, perPage, offset := pagination.ParsePageAndPerPage(r)
	filter := DecodeListTasksRequest(r)

	result, meta, err := h.usecase.List(r.Context(), userID, filter, page, perPage, offset)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Paginated(w, http.StatusOK, result, meta)
}
