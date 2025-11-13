package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codeforgood-org/golang-book-api/internal/models"
	"github.com/codeforgood-org/golang-book-api/internal/storage"
)

func TestBookHandler_HandleBooks_GET(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := NewBookHandler(store)

	// Add some test books
	store.Create(models.Book{Title: "Book 1", Author: "Author 1"})
	store.Create(models.Book{Title: "Book 2", Author: "Author 2"})

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()

	handler.HandleBooks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response models.PaginatedResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Check that data is an array
	data, ok := response.Data.([]interface{})
	if !ok {
		t.Fatal("expected data to be an array")
	}

	if len(data) != 2 {
		t.Errorf("expected 2 books, got %d", len(data))
	}

	if response.Total != 2 {
		t.Errorf("expected total 2, got %d", response.Total)
	}
}

func TestBookHandler_HandleBooks_POST(t *testing.T) {
	tests := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid book",
			payload: models.Book{
				Title:  "Test Book",
				Author: "Test Author",
			},
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "missing title",
			payload: models.Book{
				Author: "Test Author",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "missing author",
			payload: models.Book{
				Title: "Test Book",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid json",
			payload:        "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := storage.NewMemoryStorage()
			handler := NewBookHandler(store)

			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.HandleBooks(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectError {
				var errResp map[string]string
				json.NewDecoder(w.Body).Decode(&errResp)
				if _, exists := errResp["error"]; !exists {
					t.Error("expected error response to contain 'error' field")
				}
			} else {
				var book models.Book
				if err := json.NewDecoder(w.Body).Decode(&book); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if book.ID == 0 {
					t.Error("expected book to have an ID assigned")
				}
			}
		})
	}
}

func TestBookHandler_HandleBookByID_GET(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := NewBookHandler(store)

	// Create a test book
	created, _ := store.Create(models.Book{Title: "Test Book", Author: "Test Author"})

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "existing book",
			url:            fmt.Sprintf("/books/%d", created.ID),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "non-existing book",
			url:            "/books/999999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "invalid ID",
			url:            "/books/invalid",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			handler.HandleBookByID(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestBookHandler_HandleBookByID_DELETE(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := NewBookHandler(store)

	// Create a test book
	created, _ := store.Create(models.Book{Title: "Test Book", Author: "Test Author"})

	tests := []struct {
		name           string
		bookID         int
		expectedStatus int
	}{
		{
			name:           "delete existing book",
			bookID:         created.ID,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "delete non-existing book",
			bookID:         999999,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/books/%d", tt.bookID)
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			w := httptest.NewRecorder()

			handler.HandleBookByID(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestBookHandler_HandleBooks_MethodNotAllowed(t *testing.T) {
	store := storage.NewMemoryStorage()
	handler := NewBookHandler(store)

	req := httptest.NewRequest(http.MethodPut, "/books", nil)
	w := httptest.NewRecorder()

	handler.HandleBooks(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
