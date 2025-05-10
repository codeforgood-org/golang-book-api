package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books = make([]Book, 0)
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookByIDHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var b Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		mu.Lock()
		b.ID = rand.Intn(1000000)
		books = append(books, b)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, b := range books {
		if b.ID == id {
			if r.Method == http.MethodDelete {
				books = append(books[:i], books[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.NotFound(w, r)
}
