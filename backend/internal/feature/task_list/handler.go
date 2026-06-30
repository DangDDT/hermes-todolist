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

// List godoc
// @Summary      List tasks
// @Description  List tasks with filters and pagination (requires authentication)
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        status    query     string  false  "Filter by status"        Enums(TODO, IN_PROGRESS, DONE, CANCELLED)
// @Param        priority  query     string  false  "Filter by priority"      Enums(LOW, MEDIUM, HIGH, URGENT)
// @Param        search    query     string  false  "Search in title/description"
// @Param        page      query     int     false  "Page number"             default(1)
// @Param        per_page  query     int     false  "Items per page"          default(20)
// @Success      200       {object}  response.Envelope
// @Failure      400       {object}  response.Envelope
// @Failure      401       {object}  response.Envelope
// @Router       /tasks [get]
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
