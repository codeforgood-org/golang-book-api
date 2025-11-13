package models

import (
	"net/http"
	"strings"
)

// BookFilters holds filter parameters for book queries
type BookFilters struct {
	Title  string
	Author string
	Search string
}

// ParseBookFilters extracts filter parameters from request
func ParseBookFilters(r *http.Request) BookFilters {
	return BookFilters{
		Title:  strings.TrimSpace(r.URL.Query().Get("title")),
		Author: strings.TrimSpace(r.URL.Query().Get("author")),
		Search: strings.TrimSpace(r.URL.Query().Get("search")),
	}
}

// Match checks if a book matches the filters
func (f BookFilters) Match(book Book) bool {
	// If search is provided, match against title or author
	if f.Search != "" {
		searchLower := strings.ToLower(f.Search)
		titleLower := strings.ToLower(book.Title)
		authorLower := strings.ToLower(book.Author)

		if !strings.Contains(titleLower, searchLower) &&
			!strings.Contains(authorLower, searchLower) {
			return false
		}
	}

	// Match specific title filter
	if f.Title != "" {
		titleLower := strings.ToLower(book.Title)
		filterLower := strings.ToLower(f.Title)
		if !strings.Contains(titleLower, filterLower) {
			return false
		}
	}

	// Match specific author filter
	if f.Author != "" {
		authorLower := strings.ToLower(book.Author)
		filterLower := strings.ToLower(f.Author)
		if !strings.Contains(authorLower, filterLower) {
			return false
		}
	}

	return true
}

// HasFilters returns true if any filters are set
func (f BookFilters) HasFilters() bool {
	return f.Title != "" || f.Author != "" || f.Search != ""
}
