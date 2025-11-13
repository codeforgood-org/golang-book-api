package models

import "testing"

func TestBookFilters_Match(t *testing.T) {
	tests := []struct {
		name    string
		filters BookFilters
		book    Book
		want    bool
	}{
		{
			name:    "no filters - should match",
			filters: BookFilters{},
			book:    Book{Title: "Any Book", Author: "Any Author"},
			want:    true,
		},
		{
			name:    "title filter - exact match",
			filters: BookFilters{Title: "Go Programming"},
			book:    Book{Title: "Go Programming", Author: "Author"},
			want:    true,
		},
		{
			name:    "title filter - partial match",
			filters: BookFilters{Title: "Go"},
			book:    Book{Title: "The Go Programming Language", Author: "Author"},
			want:    true,
		},
		{
			name:    "title filter - case insensitive",
			filters: BookFilters{Title: "go"},
			book:    Book{Title: "Go Programming", Author: "Author"},
			want:    true,
		},
		{
			name:    "title filter - no match",
			filters: BookFilters{Title: "Python"},
			book:    Book{Title: "Go Programming", Author: "Author"},
			want:    false,
		},
		{
			name:    "author filter - match",
			filters: BookFilters{Author: "Donovan"},
			book:    Book{Title: "Book", Author: "Alan A. A. Donovan"},
			want:    true,
		},
		{
			name:    "author filter - no match",
			filters: BookFilters{Author: "Smith"},
			book:    Book{Title: "Book", Author: "Alan Donovan"},
			want:    false,
		},
		{
			name:    "search - match in title",
			filters: BookFilters{Search: "Programming"},
			book:    Book{Title: "Go Programming", Author: "Author"},
			want:    true,
		},
		{
			name:    "search - match in author",
			filters: BookFilters{Search: "Donovan"},
			book:    Book{Title: "Book", Author: "Alan Donovan"},
			want:    true,
		},
		{
			name:    "search - no match",
			filters: BookFilters{Search: "Python"},
			book:    Book{Title: "Go Programming", Author: "Donovan"},
			want:    false,
		},
		{
			name:    "multiple filters - all match",
			filters: BookFilters{Title: "Go", Author: "Donovan"},
			book:    Book{Title: "Go Programming", Author: "Alan Donovan"},
			want:    true,
		},
		{
			name:    "multiple filters - one doesn't match",
			filters: BookFilters{Title: "Python", Author: "Donovan"},
			book:    Book{Title: "Go Programming", Author: "Alan Donovan"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.filters.Match(tt.book); got != tt.want {
				t.Errorf("BookFilters.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookFilters_HasFilters(t *testing.T) {
	tests := []struct {
		name    string
		filters BookFilters
		want    bool
	}{
		{
			name:    "no filters",
			filters: BookFilters{},
			want:    false,
		},
		{
			name:    "title filter only",
			filters: BookFilters{Title: "Go"},
			want:    true,
		},
		{
			name:    "author filter only",
			filters: BookFilters{Author: "Donovan"},
			want:    true,
		},
		{
			name:    "search filter only",
			filters: BookFilters{Search: "Programming"},
			want:    true,
		},
		{
			name:    "multiple filters",
			filters: BookFilters{Title: "Go", Author: "Donovan"},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.filters.HasFilters(); got != tt.want {
				t.Errorf("BookFilters.HasFilters() = %v, want %v", got, tt.want)
			}
		})
	}
}
