package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book = []Book{}

// Book is a struct that represents a single Book
type Book struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	NoOfPages int     `json:"noofpages"`
	Price     float32 `json:"price"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")

	http.ListenAndServe(":5000", router)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)

	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// get the ID of the book from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(books) {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}

	book := books[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
