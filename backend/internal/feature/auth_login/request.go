package auth_login

import (
	"encoding/json"
	"net/http"
)

// LoginRequest is the request body for login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DecodeLoginRequest parses the JSON request body.
func DecodeLoginRequest(r *http.Request) (*LoginRequest, error) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// Validate checks the request fields.
func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return &fieldError{Field: "username", Message: "username is required"}
	}
	if r.Password == "" {
		return &fieldError{Field: "password", Message: "password is required"}
	}
	return nil
}

type fieldError struct {
	Field   string
	Message string
}

func (e *fieldError) Error() string {
	return e.Field + ": " + e.Message
}
