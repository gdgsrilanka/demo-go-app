package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	books = append(books, Book{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	books = append(books, Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"})
	books = append(books, Book{ID: "3", Title: "1984", Author: "George Orwell"})

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Folks! \n")
	})

	// /hello endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Folks! \nThis is endpoint 2\n")
	})

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w, r)
		case http.MethodPost:
			addBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /books/{id} endpoint with different methods
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 3 && parts[2] != "" {
			bookID := parts[2]
			switch r.Method {
			case http.MethodGet:
				getBookByID(w, r, bookID)
			case http.MethodPut:
				updateBook(w, r, bookID)
			case http.MethodDelete:
				deleteBook(w, r, bookID)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		} else {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
		}
	})

	fmt.Println("Starting Folks API Server...")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// Handler for GET /books
func getBooks(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler for POST /books
func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	books = append(books, newBook)

	response, err := json.Marshal(newBook)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler for GET /books/{id}
func getBookByID(w http.ResponseWriter, r *http.Request, bookID string) {
	book, err := findBookByID(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler for PUT /books/{id}
func updateBook(w http.ResponseWriter, r *http.Request, bookID string) {
	_, index, err := findBookByIDWithIndex(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	books[index] = updatedBook

	response, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Handler for DELETE /books/{id}
func deleteBook(w http.ResponseWriter, r *http.Request, bookID string) {
	_, index, err := findBookByIDWithIndex(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	deletedBook := books[index]
	books = append(books[:index], books[index+1:]...)

	response, err := json.Marshal(map[string]string{"message": "Book deleted", "id": deletedBook.ID, "title": deletedBook.Title, "author": deletedBook.Author})
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func findBookByID(bookID string) (Book, error) {
	for _, book := range books {
		if book.ID == bookID {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("Book with ID %s not found", bookID)
}

func findBookByIDWithIndex(bookID string) (Book, int, error) {
	for i, book := range books {
		if book.ID == bookID {
			return book, i, nil
		}
	}
	return Book{}, -1, fmt.Errorf("Book with ID %s not found", bookID)
}
