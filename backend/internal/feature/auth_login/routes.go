package auth_login

import (
	"github.com/go-chi/chi/v5"
)

// Routes returns the router for the auth/login feature.
func Routes(usecase *Usecase) chi.Router {
	r := chi.NewRouter()
	h := NewHandler(usecase)
	r.Post("/", h.Login)
	return r
}
