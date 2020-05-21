package main

import (
	"encoding/json"
	"github.com/rlr524/restapi/src/github.com/gorilla/mux"
	"net/http"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get the params from the entered url
	// loop through books and find the one with the id from params
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
