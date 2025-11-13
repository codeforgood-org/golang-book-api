package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeforgood-org/golang-book-api/internal/models"
)

var sampleBooks = []models.Book{
	{Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan"},
	{Title: "Clean Code", Author: "Robert C. Martin"},
	{Title: "Design Patterns", Author: "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides"},
	{Title: "The Pragmatic Programmer", Author: "Andrew Hunt and David Thomas"},
	{Title: "Introduction to Algorithms", Author: "Thomas H. Cormen"},
	{Title: "Code Complete", Author: "Steve McConnell"},
	{Title: "Refactoring", Author: "Martin Fowler"},
	{Title: "The Clean Coder", Author: "Robert C. Martin"},
	{Title: "Head First Design Patterns", Author: "Eric Freeman and Elisabeth Robson"},
	{Title: "You Don't Know JS", Author: "Kyle Simpson"},
	{Title: "Eloquent JavaScript", Author: "Marijn Haverbeke"},
	{Title: "JavaScript: The Good Parts", Author: "Douglas Crockford"},
	{Title: "Python Crash Course", Author: "Eric Matthes"},
	{Title: "Effective Java", Author: "Joshua Bloch"},
	{Title: "Clean Architecture", Author: "Robert C. Martin"},
	{Title: "Domain-Driven Design", Author: "Eric Evans"},
	{Title: "Microservices Patterns", Author: "Chris Richardson"},
	{Title: "Building Microservices", Author: "Sam Newman"},
	{Title: "Site Reliability Engineering", Author: "Betsy Beyer, Chris Jones, Jennifer Petoff, Niall Richard Murphy"},
	{Title: "The DevOps Handbook", Author: "Gene Kim, Jez Humble, Patrick Debois, John Willis"},
}

func main() {
	baseURL := os.Getenv("API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	fmt.Printf("Seeding data to %s/books\n", baseURL)

	for i, book := range sampleBooks {
		data, err := json.Marshal(book)
		if err != nil {
			log.Printf("Failed to marshal book %d: %v", i+1, err)
			continue
		}

		resp, err := http.Post(
			fmt.Sprintf("%s/books", baseURL),
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			log.Printf("Failed to create book %d: %v", i+1, err)
			continue
		}

		if resp.StatusCode != http.StatusCreated {
			log.Printf("Failed to create book %d: status %d", i+1, resp.StatusCode)
			resp.Body.Close()
			continue
		}

		var createdBook models.Book
		if err := json.NewDecoder(resp.Body).Decode(&createdBook); err != nil {
			log.Printf("Failed to decode response for book %d: %v", i+1, err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		fmt.Printf("✓ Created: %s (ID: %d)\n", createdBook.Title, createdBook.ID)
	}

	fmt.Printf("\nSeeded %d books successfully!\n", len(sampleBooks))
}
