package library_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/library/mock"
)

func TestHandler_getBooks(t *testing.T) {
	service := mock.NewMockService()
	router := chi.NewRouter()

	usr := mock.NewMockUserServiceClient()

	// Handler creation
	h := library.NewHandler(router, service, usr)
	h.Register()

	// Create GET request to get books list
	req, err := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create ResponseRecorder for testing
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status
	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusOK, rr.Code)
		}
	})

	// Check the answer
	t.Run("body", func(t *testing.T) {
		var got []library.Book
		err := json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		want := []library.Book{
			{
				ID:          "1",
				Title:       "Book One",
				Author:      "Author One",
				Description: "Description One",
			},
			{
				ID:          "2",
				Title:       "Book Two",
				Author:      "Author Two",
				Description: "Description Two",
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("GET /api/v1/books mismatch: (-want +got)\n%s", diff)
		}
	})
}

func TestHandler_getBookByID(t *testing.T) {
	service := mock.NewMockService()
	router := chi.NewRouter()

	usr := mock.NewMockUserServiceClient()

	// Handler creation
	h := library.NewHandler(router, service, usr)
	h.Register()

	// Create GET request to get book with ID 1
	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create ResponseRecorder for testing
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status
	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusOK, rr.Code)
		}
	})

	// Check the answer
	t.Run("body", func(t *testing.T) {
		var got library.Book
		err := json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		want := library.Book{
			ID:          "1",
			Title:       "Book One",
			Author:      "Author One",
			Description: "Description One",
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("GET /api/v1/books/1 mismatch: (-want +got)\n%s", diff)
		}
	})
}

func TestHandler_createBook(t *testing.T) {
	service := mock.NewMockService()
	router := chi.NewRouter()

	usr := mock.NewMockUserServiceClient()

	// Handler creation
	h := library.NewHandler(router, service, usr)
	h.Register()

	// Create POST request to create new book
	newBook := `{"title": "New Book", "author": "New Author", "description": "New Description"}`
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books/new", strings.NewReader(newBook))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "No matter")

	// Create ResponseRecorder for testing
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status
	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusCreated {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	// Check the answer
	t.Run("body", func(t *testing.T) {
		var got map[string]string
		err := json.NewDecoder(rr.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		want := map[string]string{"id": "3"}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("POST /api/v1/books mismatch: (-want +got)\n%s", diff)
		}
	})
}

func TestHandler_updateBook(t *testing.T) {
	service := mock.NewMockService()
	router := chi.NewRouter()

	usr := mock.NewMockUserServiceClient()

	// Handler creation
	h := library.NewHandler(router, service, usr)
	h.Register()

	// Create PUT request to update book
	updatedBook := `{"title": "Updated Book", "author": "Updated Author", "description": "Updated Description"}`
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books/1", strings.NewReader(updatedBook))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "No matter what")

	// Create ResponseRecorder for testing
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status
	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusNoContent, rr.Code)
		}
	})
}

func TestHandler_deleteBook(t *testing.T) {
	service := mock.NewMockService()
	router := chi.NewRouter()

	usr := mock.NewMockUserServiceClient()

	// Handler creation
	h := library.NewHandler(router, service, usr)
	h.Register()

	// Create DELETE request to delete book
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Now matter what")

	// Create ResponseRecorder for testing
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the answer
	t.Run("status", func(t *testing.T) {
		if rr.Code != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusNoContent, rr.Code)
		}
	})
}
