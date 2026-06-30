package task_get

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the task get feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Get("/{id}", h.Get)
	return r
}
