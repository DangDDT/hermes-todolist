package auth_register

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the auth/register feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Post("/", h.Register)
	return r
}
