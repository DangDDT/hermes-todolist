package task_list

import (
	"encoding/json"
	"net/http"
)

// ListTasksRequest is the parsed query parameters for listing tasks.
type ListTasksRequest struct {
	Status   string
	Priority string
	Search   string
}

// DecodeListTasksRequest parses query params.
func DecodeListTasksRequest(r *http.Request) *ListTasksRequest {
	return &ListTasksRequest{
		Status:   r.URL.Query().Get("status"),
		Priority: r.URL.Query().Get("priority"),
		Search:   r.URL.Query().Get("search"),
	}
}

// Ensure json import is used.
var _ = json.Encoder{}
