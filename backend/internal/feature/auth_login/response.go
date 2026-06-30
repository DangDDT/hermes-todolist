package auth_login

import (
	"time"
)

// LoginResponse is the response for successful login.
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	User      UserInfo `json:"user"`
}

// UserInfo contains public user data.
type UserInfo struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
}
