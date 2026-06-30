package auth_register

import (
	"encoding/json"
	"net/http"

)

// RegisterUserRequest is the request body for user registration.
type RegisterUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
}

// DecodeRegisterUserRequest parses the JSON request body.
func DecodeRegisterUserRequest(r *http.Request) (*RegisterUserRequest, error) {
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// Validate checks the request fields.
func (r *RegisterUserRequest) Validate() error {
	if r.Username == "" {
		return fieldError("username", "username is required")
	}
	if r.Password == "" {
		return fieldError("password", "password is required")
	}
	if len(r.Password) < 8 {
		return fieldError("password", "password must be at least 8 characters")
	}
	return nil
}

func fieldError(field, message string) error {
	return &FieldError{Field: field, Message: message}
}

// FieldError represents a validation error on a specific field.
type FieldError struct {
	Field   string
	Message string
}

func (e *FieldError) Error() string {
	return e.Field + ": " + e.Message
}
