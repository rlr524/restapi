package main

import (
	"github.com/rlr524/restapi/src/github.com/gorilla/mux"
	"log"
	"net/http"
)

// book structs (model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

// init books var as a slice Book struct
var books []Book

func main() {
	// init mux router
	r := mux.NewRouter()

	// mock book data TODO: implement database
	books = append(books, Book{ID: "1", Isbn: "0596007124", Title: "Head First Design Patterns",
		Author: &Author{Fname: "Eric", Lname: "Freeman"}})
	books = append(books, Book{ID: "2", Isbn: "0134663705", Title: "Learning Node.js",
		Author: &Author{Fname: "Marc", Lname: "Wandschneider"}})

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run our server
	log.Fatal(http.ListenAndServe(":8000", r))
}
