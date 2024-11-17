package library

import "context"

// Book structure represents book entity
type Book struct {
	ID          string
	Title       string
	Author      string
	Description string
}

// BookService defines the interface for interacting with books (business logic)
type BookService interface {
	GetBooks(ctx context.Context, criteria string) ([]Book, error)
	GetBookByID(ctx context.Context, id string) (*Book, error)
	CreateBook(ctx context.Context, book Book) (string, error)
	UpdateBook(ctx context.Context, id string, book Book) error
	DeleteBook(ctx context.Context, id string) error
}

// BookStore defines the inteface for database interactions related to books
type BookStore interface {
	LoadBooks(ctx context.Context, criteria string) ([]Book, error)
	LoadBookByID(ctx context.Context, id string) (*Book, error)
	SaveBook(ctx context.Context, book Book) (string, error)
	UpdateBook(ctx context.Context, id string, book Book) error
	DeleteBook(ctx context.Context, id string) error
}
