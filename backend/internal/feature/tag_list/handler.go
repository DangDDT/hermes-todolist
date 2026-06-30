package tag_list

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

// TagItem represents a tag in the list response.
type TagItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListTagsResponse is the response for the tag list endpoint.
type ListTagsResponse struct {
	Tags []TagItem `json:"tags"`
}

// Usecase handles listing tags.
type Usecase struct {
}

// NewUsecase creates a new tag list usecase.
func NewUsecase() *Usecase {
	return &Usecase{}
}

// List retrieves all tags.
func (uc *Usecase) List(ctx context.Context) (*ListTagsResponse, error) {
	// Stub: return empty list for now.
	return &ListTagsResponse{Tags: []TagItem{}}, nil
}

// Handler handles GET /tags.
type Handler struct {
	usecase *Usecase
}

// NewHandler creates a new tag list handler.
func NewHandler(usecase *Usecase) *Handler {
	return &Handler{usecase: usecase}
}

// List handles the list tags request.
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.List(r.Context())
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, http.StatusOK, result, nil)
}

// Routes returns the router for the tag list feature.
func Routes() chi.Router {
	r := chi.NewRouter()
	h := NewHandler(NewUsecase())
	r.Get("/", h.List)
	return r
}
