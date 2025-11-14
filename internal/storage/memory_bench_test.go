package storage

import (
	"testing"

	"github.com/codeforgood-org/golang-book-api/internal/models"
)

func BenchmarkMemoryStorage_Create(b *testing.B) {
	storage := NewMemoryStorage()
	book := models.Book{
		Title:  "Benchmark Book",
		Author: "Benchmark Author",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.Create(book)
	}
}

func BenchmarkMemoryStorage_GetAll(b *testing.B) {
	storage := NewMemoryStorage()

	// Pre-populate with books
	for i := 0; i < 100; i++ {
		storage.Create(models.Book{
			Title:  "Book",
			Author: "Author",
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.GetAll()
	}
}

func BenchmarkMemoryStorage_GetByID(b *testing.B) {
	storage := NewMemoryStorage()

	// Create a book to search for
	book, _ := storage.Create(models.Book{
		Title:  "Benchmark Book",
		Author: "Benchmark Author",
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.GetByID(book.ID)
	}
}

func BenchmarkMemoryStorage_Update(b *testing.B) {
	storage := NewMemoryStorage()

	// Create a book to update
	book, _ := storage.Create(models.Book{
		Title:  "Original Title",
		Author: "Original Author",
	})

	updatedBook := models.Book{
		Title:  "Updated Title",
		Author: "Updated Author",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		storage.Update(book.ID, updatedBook)
	}
}

func BenchmarkMemoryStorage_Delete(b *testing.B) {
	storage := NewMemoryStorage()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// Create a book for each iteration
		book, _ := storage.Create(models.Book{
			Title:  "Book to Delete",
			Author: "Author",
		})
		b.StartTimer()

		storage.Delete(book.ID)
	}
}

func BenchmarkMemoryStorage_ConcurrentReads(b *testing.B) {
	storage := NewMemoryStorage()

	// Pre-populate
	for i := 0; i < 100; i++ {
		storage.Create(models.Book{
			Title:  "Book",
			Author: "Author",
		})
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			storage.GetAll()
		}
	})
}

func BenchmarkMemoryStorage_ConcurrentWrites(b *testing.B) {
	storage := NewMemoryStorage()
	book := models.Book{
		Title:  "Concurrent Book",
		Author: "Concurrent Author",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			storage.Create(book)
		}
	})
}
