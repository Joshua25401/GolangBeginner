package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct Model
type Book struct {
	Book_ID     string  `json:"book_id"`
	Book_Isbn   string  `json:"book_isbn"`
	Book_Title  string  `json:"book_title"`
	Book_Author *Author `json:"book_author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books data as a Slice
var books []Book

// Function to handle API requests
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Origin", "*")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Origin", "*")

	params := mux.Vars(r)

	for _, item := range books {
		if item.Book_ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Return empty if not found
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Origin", "*")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatal(err)
		return
	}

	book.Book_ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Book_ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	// Router
	router := mux.NewRouter()

	// Pre-Populate data
	books = append(books,
		Book{
			Book_ID:    "1",
			Book_Isbn:  "11319029",
			Book_Title: "Dasar Pemrograman",
			Book_Author: &Author{
				Firstname: "Joshua",
				Lastname:  "Ryandafres"}},
		Book{
			Book_ID:    "2",
			Book_Isbn:  "11319014",
			Book_Title: "Elasticsearch Document",
			Book_Author: &Author{
				Firstname: "Alex Sander",
				Lastname:  "Hutapea",
			}},
	)
	log.Println("API Server is Running...")

	// Handle Routes
	router.HandleFunc("/api/books", getBooks).Methods("GET")     // Endpoint : http://localhost:8000/api/books
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET") // Endpoint : http://localhost:8000/api/books/2
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
