package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/codeforgood-org/golang-book-api/internal/models"
	"github.com/codeforgood-org/golang-book-api/internal/storage"
	"github.com/codeforgood-org/golang-book-api/pkg/logger"
)

// BookHandler handles book-related HTTP requests
type BookHandler struct {
	storage storage.Storage
}

// NewBookHandler creates a new book handler
func NewBookHandler(storage storage.Storage) *BookHandler {
	return &BookHandler{
		storage: storage,
	}
}

// HandleBooks handles requests to /books endpoint
func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getBooks(w, r)
	case http.MethodPost:
		h.createBook(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// HandleBookByID handles requests to /books/{id} endpoint
func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, models.ErrInvalidID.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getBookByID(w, r, id)
	case http.MethodDelete:
		h.deleteBook(w, r, id)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// getBooks returns all books
func (h *BookHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.storage.GetAll()
	if err != nil {
		logger.Error.Printf("Failed to get books: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

// createBook creates a new book
func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate the book
	if err := book.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create the book
	createdBook, err := h.storage.Create(book)
	if err != nil {
		logger.Error.Printf("Failed to create book: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	respondWithJSON(w, http.StatusCreated, createdBook)
}

// getBookByID returns a book by ID
func (h *BookHandler) getBookByID(w http.ResponseWriter, r *http.Request, id int) {
	book, err := h.storage.GetByID(id)
	if err != nil {
		if err == models.ErrBookNotFound {
			respondWithError(w, http.StatusNotFound, "Book not found")
			return
		}
		logger.Error.Printf("Failed to get book: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve book")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

// deleteBook deletes a book by ID
func (h *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request, id int) {
	err := h.storage.Delete(id)
	if err != nil {
		if err == models.ErrBookNotFound {
			respondWithError(w, http.StatusNotFound, "Book not found")
			return
		}
		logger.Error.Printf("Failed to delete book: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// respondWithJSON writes a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error.Printf("Failed to encode response: %v", err)
	}
}

// respondWithError writes an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
