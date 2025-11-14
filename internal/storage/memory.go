package storage

import (
	"math/rand"
	"sync"
	"time"

	"github.com/codeforgood-org/golang-book-api/internal/models"
)

// MemoryStorage implements in-memory storage for books
type MemoryStorage struct {
	books []models.Book
	mu    sync.RWMutex
	rng   *rand.Rand
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		books: make([]models.Book, 0),
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetAll returns all books
func (s *MemoryStorage) GetAll() ([]models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to prevent external modifications
	booksCopy := make([]models.Book, len(s.books))
	copy(booksCopy, s.books)
	return booksCopy, nil
}

// GetByID returns a book by its ID
func (s *MemoryStorage) GetByID(id int) (*models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, book := range s.books {
		if book.ID == id {
			// Return a copy
			bookCopy := book
			return &bookCopy, nil
		}
	}
	return nil, models.ErrBookNotFound
}

// Create adds a new book and returns it with an assigned ID
func (s *MemoryStorage) Create(book models.Book) (*models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Generate a unique ID
	book.ID = s.rng.Intn(1000000)
	s.books = append(s.books, book)

	return &book, nil
}

// Update updates an existing book
func (s *MemoryStorage) Update(id int, book models.Book) (*models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, b := range s.books {
		if b.ID == id {
			// Preserve the original ID
			book.ID = id
			s.books[i] = book
			return &book, nil
		}
	}
	return nil, models.ErrBookNotFound
}

// Delete removes a book by its ID
func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return models.ErrBookNotFound
}
