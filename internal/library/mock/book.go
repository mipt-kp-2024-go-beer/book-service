package mock

import (
	"context"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
)

// MockBookService implements the library.BookService interface for testing purposes
type Mock struct{}

func NewMockService() *Mock {
	return &Mock{}
}

func (m *Mock) GetBooks(ctx context.Context, criteria string) ([]library.Book, error) {
	return []library.Book{
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
	}, nil
}

// GetBookByID mocks the GetBookByID method from the BookService interface
func (m *Mock) GetBookByID(ctx context.Context, id string) (*library.Book, error) {
	if id == "1" {
		return &library.Book{
			ID:          "1",
			Title:       "Book One",
			Author:      "Author One",
			Description: "Description One",
		}, nil
	}
	return nil, nil
}

func (m *Mock) CreateBook(ctx context.Context, book library.Book) (string, error) {
	return "3", nil
}

// UpdateBook mocks the UpdateBook method from the BookService interface
func (m *Mock) UpdateBook(ctx context.Context, id string, book library.Book) error {
	return nil
}

// DeleteBook mocks the DeleteBook method from the BookService interface
func (m *Mock) DeleteBook(ctx context.Context, id string) error {
	return nil
}
