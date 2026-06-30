package apperrors

import (
	"fmt"
	"net/http"
)

// AppError represents an application-level error with HTTP metadata.
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error.
func (e *AppError) Unwrap() error {
	return e.Err
}

// NotFound creates a 404 error.
func NotFound(entity string, err error) *AppError {
	return &AppError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s not found", entity),
		HTTPStatus: http.StatusNotFound,
		Err:        err,
	}
}

// ValidationError creates a 422 error.
func ValidationError(message string, err error) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		HTTPStatus: http.StatusUnprocessableEntity,
		Err:        err,
	}
}

// Unauthorized creates a 401 error.
func Unauthorized(message string, err error) *AppError {
	if message == "" {
		message = "unauthorized"
	}
	return &AppError{
		Code:       "UNAUTHORIZED",
		Message:    message,
		HTTPStatus: http.StatusUnauthorized,
		Err:        err,
	}
}

// Conflict creates a 409 error.
func Conflict(entity, field string, err error) *AppError {
	return &AppError{
		Code:       "CONFLICT",
		Message:    fmt.Sprintf("%s with this %s already exists", entity, field),
		HTTPStatus: http.StatusConflict,
		Err:        err,
	}
}

// Internal creates a 500 error.
func Internal(err error) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    "An internal error occurred",
		HTTPStatus: http.StatusInternalServerError,
		Err:        err,
	}
}

// Forbidden creates a 403 error.
func Forbidden(message string, err error) *AppError {
	if message == "" {
		message = "forbidden"
	}
	return &AppError{
		Code:       "FORBIDDEN",
		Message:    message,
		HTTPStatus: http.StatusForbidden,
		Err:        err,
	}
}
