package task_update

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the task update feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Put("/{id}", h.Update)
	return r
}
