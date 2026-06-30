package task_delete

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the task delete feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Delete("/{id}", h.Delete)
	return r
}
