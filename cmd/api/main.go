package main

import (
	"fmt"
	"net/http"

	"github.com/codeforgood-org/golang-book-api/internal/config"
	"github.com/codeforgood-org/golang-book-api/internal/handlers"
	"github.com/codeforgood-org/golang-book-api/internal/middleware"
	"github.com/codeforgood-org/golang-book-api/internal/storage"
	"github.com/codeforgood-org/golang-book-api/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize storage
	bookStorage := storage.NewMemoryStorage()

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookStorage)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/books", bookHandler.HandleBooks)
	mux.HandleFunc("/books/", bookHandler.HandleBookByID)

	// Apply middleware
	handler := middleware.Recovery(
		middleware.Logger(
			middleware.CORS(mux),
		),
	)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	logger.Info.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Error.Fatalf("Server failed to start: %v", err)
	}
}
