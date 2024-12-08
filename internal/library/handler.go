package library

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	router  *chi.Mux
	service BookService
	userSVC UserService
}

func NewHandler(router *chi.Mux, service BookService, userSVC UserService) *Handler {
	return &Handler{
		router:  router,
		service: service,
		userSVC: userSVC,
	}
}

// Register routes for the Handler
func (h *Handler) Register() {
	h.router.Group(func(r chi.Router) {
		r.Get("/api/v1/books", h.getBooks)
		r.Get("/api/v1/books/{id}", h.getBookByID)
		r.Post("/api/v1/books/new", h.createBook)
		r.Post("/api/v1/books/{id}", h.updateBook)
		r.Delete("/api/v1/books/{id}", h.deleteBook)
	})
}

// Handles GET request to fetch all books
func (h *Handler) getBooks(w http.ResponseWriter, r *http.Request) {
	criteria := r.URL.Query().Get("criteria")
	ctx := r.Context()

	// Get list of books from the service
	books, err := h.service.GetBooks(ctx, criteria)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get books: %v", err), http.StatusInternalServerError)
		return
	}

	// Return books as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode books: %v", err), http.StatusInternalServerError)
		return
	}
}

// Handles GET request to fetch a single book by ID
func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	// Get the book by ID from the service
	book, err := h.service.GetBookByID(ctx, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get book by ID: %v", err), http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return the book as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode book: %v", err), http.StatusInternalServerError)
		return
	}
}

// Handles POST requests to create a new book
func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	// First, let's get authorization token
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Request to 'user' microservice to get permissions
	manage, err := h.userSVC.CheckPermissions(token, PermManageBooks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking permission: %v", err), http.StatusInternalServerError)
		return
	}

	if !manage {
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	// Create the book via the service
	id, err := h.service.CreateBook(ctx, book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create book: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the created book's ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"id": id}); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

// Handles PUT request to update a book by ID
func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	// First, let's get authorization token
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Request to 'user' microservice to get permissions
	manage, err := h.userSVC.CheckPermissions(token, PermManageBooks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking permission: %v", err), http.StatusInternalServerError)
		return
	}

	if !manage {
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
		return
	}

	id := chi.URLParam(r, "id")
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	// Update the book via the service
	if err := h.service.UpdateBook(ctx, id, book); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update book: %v", err), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusNoContent)
}

// Handles DELETE request to delete a book by ID
func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	// First, let's get authorization token
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Request to 'user' microservice to get permissions
	manage, err := h.userSVC.CheckPermissions(token, PermManageBooks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking permission: %v", err), http.StatusInternalServerError)
		return
	}

	if !manage {
		http.Error(w, "Insufficient permissions", http.StatusForbidden)
		return
	}

	id := chi.URLParam(r, "id")
	ctx := r.Context()

	// Delete the book via the service
	if err := h.service.DeleteBook(ctx, id); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete book: %v", err), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusNoContent)
}
