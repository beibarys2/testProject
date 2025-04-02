package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testProject/bookstore/models"
)

// In-memory storage
var books = []models.Book{}
var lastID = 0

// Get all books with optional filters and pagination
func GetBooks(w http.ResponseWriter, r *http.Request) {
	categoryFilter := r.URL.Query().Get("category")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	// Фильтрация по категории
	var filteredBooks []models.Book
	for _, book := range books {
		if categoryFilter == "" || strings.EqualFold(categoryFilter, strconv.Itoa(book.CategoryID)) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	// Пагинация (5 книг на страницу)
	start := (page - 1) * 5
	end := start + 5
	if start >= len(filteredBooks) {
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	json.NewEncoder(w).Encode(filteredBooks[start:end])
}

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Simple validation
	if newBook.Title == "" || newBook.AuthorID == 0 || newBook.CategoryID == 0 || newBook.Price <= 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	lastID++
	newBook.ID = lastID
	books = append(books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// Update an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	bookID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == bookID {
			// Обновляем поля
			books[i].Title = updatedBook.Title
			books[i].AuthorID = updatedBook.AuthorID
			books[i].CategoryID = updatedBook.CategoryID
			books[i].Price = updatedBook.Price

			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	bookID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == bookID {
			// Удаляем книгу из списка
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
