package library

import (
	"context"

	"github.com/pkg/errors"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/oops"
)

type AppBookService struct {
	store BookStore
}

func NewBookService(s BookStore) *AppBookService {
	return &AppBookService{store: s}
}

func (s *AppBookService) GetBooks(ctx context.Context, criteria string) ([]Book, error) {
	// Fetch books from store (database)
	books, err := s.store.LoadBooks(ctx, criteria)
	if err != nil {
		return nil, errors.Wrap(err, oops.ErrLoadBooks.Error())
	}
	return books, nil
}

func (s *AppBookService) CreateBook(ctx context.Context, book Book) (string, error) {
	// Save book in the store (database)
	id, err := s.store.SaveBook(ctx, book)
	if err != nil {
		return "", errors.Wrap(err, oops.ErrCreateBook.Error())
	}
	return id, nil
}

func (s *AppBookService) GetBookByID(ctx context.Context, id string) (*Book, error) {
	// Fetch a single book by ID
	book, err := s.store.LoadBookByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, oops.ErrLoadBooks.Error())
	}
	return book, nil
}

func (s *AppBookService) UpdateBook(ctx context.Context, id string, book Book) error {
	// Update the book in the store
	err := s.store.UpdateBook(ctx, id, book)
	if err != nil {
		return errors.Wrap(err, oops.ErrUpdateBook.Error())
	}
	return nil
}

func (s *AppBookService) DeleteBook(ctx context.Context, id string) error {
	// Delete book from the store
	err := s.store.DeleteBook(ctx, id)
	if err != nil {
		return errors.Wrap(err, oops.ErrDeleteBook.Error())
	}
	return nil
}
