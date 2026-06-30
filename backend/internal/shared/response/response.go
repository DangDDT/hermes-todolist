package response

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/DangDDT/hermes-todolist/backend/internal/shared/apperrors"
)

// Envelope is the standard API response envelope.
type Envelope struct {
	Data  any    `json:"data,omitempty"`
	Meta  any    `json:"meta,omitempty"`
	Error any    `json:"error,omitempty"`
	Code  string `json:"code,omitempty"`
}

// JSON writes a JSON response with the given status and payload.
func JSON(w http.ResponseWriter, status int, data any) {
	writeJSON(w, status, Envelope{Data: data})
}

// Success writes a successful JSON response with data and optional metadata.
func Success(w http.ResponseWriter, status int, data any, meta any) {
	writeJSON(w, status, Envelope{Data: data, Meta: meta})
}

// Error writes an error response based on an AppError.
func Error(w http.ResponseWriter, err error) {
	appErr, ok := err.(*apperrors.AppError)
	if !ok {
		slog.Error("unexpected non-AppError passed to response.Error", "error", err)
		appErr = apperrors.Internal(err)
	}
	writeJSON(w, appErr.HTTPStatus, Envelope{
		Code:  appErr.Code,
		Error: appErr.Message,
	})
}

// Paginated write a JSON response with paginated data.
func Paginated(w http.ResponseWriter, status int, data any, meta any) {
	writeJSON(w, status, Envelope{Data: data, Meta: meta})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("failed to encode JSON response", "error", err)
	}
}
