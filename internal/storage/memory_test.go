package storage

import (
	"testing"

	"github.com/codeforgood-org/golang-book-api/internal/models"
)

func TestMemoryStorage_Create(t *testing.T) {
	storage := NewMemoryStorage()

	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}

	created, err := storage.Create(book)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if created.ID == 0 {
		t.Error("expected book to have an ID assigned")
	}

	if created.Title != book.Title {
		t.Errorf("expected title %s, got %s", book.Title, created.Title)
	}

	if created.Author != book.Author {
		t.Errorf("expected author %s, got %s", book.Author, created.Author)
	}
}

func TestMemoryStorage_GetAll(t *testing.T) {
	storage := NewMemoryStorage()

	// Initially empty
	books, err := storage.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(books) != 0 {
		t.Errorf("expected 0 books, got %d", len(books))
	}

	// Add some books
	storage.Create(models.Book{Title: "Book 1", Author: "Author 1"})
	storage.Create(models.Book{Title: "Book 2", Author: "Author 2"})

	books, err = storage.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}
}

func TestMemoryStorage_GetByID(t *testing.T) {
	storage := NewMemoryStorage()

	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}

	created, _ := storage.Create(book)

	// Get existing book
	found, err := storage.GetByID(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, found.ID)
	}

	// Get non-existing book
	_, err = storage.GetByID(999999)
	if err != models.ErrBookNotFound {
		t.Errorf("expected ErrBookNotFound, got %v", err)
	}
}

func TestMemoryStorage_Delete(t *testing.T) {
	storage := NewMemoryStorage()

	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
	}

	created, _ := storage.Create(book)

	// Delete existing book
	err := storage.Delete(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify deletion
	_, err = storage.GetByID(created.ID)
	if err != models.ErrBookNotFound {
		t.Errorf("expected ErrBookNotFound after deletion, got %v", err)
	}

	// Delete non-existing book
	err = storage.Delete(999999)
	if err != models.ErrBookNotFound {
		t.Errorf("expected ErrBookNotFound, got %v", err)
	}
}

func TestMemoryStorage_Concurrent(t *testing.T) {
	storage := NewMemoryStorage()

	// Test concurrent writes
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n int) {
			book := models.Book{
				Title:  "Concurrent Book",
				Author: "Test Author",
			}
			storage.Create(book)
			done <- true
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	books, _ := storage.GetAll()
	if len(books) != 10 {
		t.Errorf("expected 10 books after concurrent writes, got %d", len(books))
	}
}
