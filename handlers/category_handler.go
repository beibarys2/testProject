package handlers

import (
	"encoding/json"
	"net/http"

	"testProject/bookstore/models"
)

// In-memory storage for categories
var categories = []models.Category{}
var lastCategoryID = 0

// Get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categories)
}

// Create a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if newCategory.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	lastCategoryID++
	newCategory.ID = lastCategoryID
	categories = append(categories, newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}
