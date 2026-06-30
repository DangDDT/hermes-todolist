package task_list

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the task list feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Get("/", h.List)
	return r
}
