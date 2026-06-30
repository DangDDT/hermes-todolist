package pagination

import (
	"net/http"
	"strconv"
)

const (
	DefaultPage    = 1
	DefaultPerPage = 20
	MaxPerPage     = 100
)

// Meta holds pagination metadata for list responses.
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// NewMeta creates pagination metadata.
func NewMeta(page, perPage, total int) Meta {
	totalPages := total / perPage
	if total%perPage > 0 {
		totalPages++
	}
	if totalPages < 1 {
		totalPages = 1
	}
	return Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
}

// ParsePageAndPerPage extracts page and perPage from the request query string.
func ParsePageAndPerPage(r *http.Request) (page, perPage, offset int) {
	page = parseIntParam(r, "page", DefaultPage)
	perPage = parseIntParam(r, "per_page", DefaultPerPage)

	if page < 1 {
		page = DefaultPage
	}
	if perPage < 1 {
		perPage = DefaultPerPage
	}
	if perPage > MaxPerPage {
		perPage = MaxPerPage
	}

	offset = (page - 1) * perPage
	return
}

func parseIntParam(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return n
}
