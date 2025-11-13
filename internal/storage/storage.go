package storage

import "github.com/codeforgood-org/golang-book-api/internal/models"

// Storage defines the interface for book storage operations
type Storage interface {
	// GetAll returns all books
	GetAll() ([]models.Book, error)

	// GetByID returns a book by its ID
	GetByID(id int) (*models.Book, error)

	// Create adds a new book and returns it with an assigned ID
	Create(book models.Book) (*models.Book, error)

	// Delete removes a book by its ID
	Delete(id int) error
}
