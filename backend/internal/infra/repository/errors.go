package repository

import "errors"

// Common repository errors.
var (
	ErrNotFound      = errors.New("resource not found")
	ErrNotImplemented = errors.New("not implemented yet")
)
