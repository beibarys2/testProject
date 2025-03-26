package handlers

import (
	"encoding/json"
	"net/http"
	"testProject/bookstore/models"
)

// In-memory storage for authors
var authors = []models.Author{}
var lastAuthorID = 0

// Get all authors
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authors)
}

// Create a new author
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var newAuthor models.Author
	if err := json.NewDecoder(r.Body).Decode(&newAuthor); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if newAuthor.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	lastAuthorID++
	newAuthor.ID = lastAuthorID
	authors = append(authors, newAuthor)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAuthor)
}
