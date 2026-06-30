package task_create

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the task create feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Post("/", h.Create)
	return r
}
