package auth_register

import (
	"time"
)

// RegisterUserResponse is the response for successful registration.
type RegisterUserResponse struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
}
