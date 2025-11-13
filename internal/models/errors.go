package models

import "errors"

var (
	// ErrBookNotFound is returned when a book is not found
	ErrBookNotFound = errors.New("book not found")

	// ErrInvalidTitle is returned when book title is empty
	ErrInvalidTitle = errors.New("book title cannot be empty")

	// ErrInvalidAuthor is returned when book author is empty
	ErrInvalidAuthor = errors.New("book author cannot be empty")

	// ErrInvalidID is returned when book ID is invalid
	ErrInvalidID = errors.New("invalid book ID")
)
