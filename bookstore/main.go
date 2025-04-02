package main

import (
	"net/http"
	"testProject/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Book routes
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books", handlers.DeleteBook).Methods("DELETE")

	// Author routes
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")

	// Category routes
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	http.ListenAndServe(":8080", r)
}
