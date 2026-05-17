package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books atau Add new book (Berdasarkan Method)
func handleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(books)

	case "POST":
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID
		books = append(books, book)
		json.NewEncoder(w).Encode(book)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Get, Update, atau Delete single book berdasarkan ID
func handleBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Mengambil parameter {id} menggunakan bawaan Go (Go 1.22+)
	idParam := r.PathValue("id")

	switch r.Method {
	case "GET":
		for _, item := range books {
			if item.ID == idParam {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&Book{})

	case "PUT":
		for index, item := range books {
			if item.ID == idParam {
				books = append(books[:index], books[index+1:]...)
				var book Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID = idParam
				books = append(books, book)
				json.NewEncoder(w).Encode(book)
				return
			}
		}

	case "DELETE":
		for index, item := range books {
			if item.ID == idParam {
				books = append(books[:index], books[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(books)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Main function
func main() {
	// Menggunakan ServeMux bawaan Go
	mux := http.NewServeMux()

	// Hardcoded data
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles menggunakan sintaks Go modern (Method + Path)
	mux.HandleFunc("GET /books", handleBooks)
	mux.HandleFunc("POST /books", handleBooks)

	mux.HandleFunc("GET /books/{id}", handleBookByID)
	mux.HandleFunc("PUT /books/{id}", handleBookByID)
	mux.HandleFunc("DELETE /books/{id}", handleBookByID)

	// Start server
	log.Println("Server jalan di port 8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
