package task_comment

import "github.com/go-chi/chi/v5"

// Routes returns the router for task comment operations.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Get("/", h.List)
	r.Post("/", h.Create)
	return r
}
