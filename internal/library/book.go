package library

import "context"

// Book structure represents book entity
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Stock       string `json:"stock"`
}

// Intercommunication with 'user' microservice (permission checks)
type UserService interface {
	CheckPermissions(token string, mask uint) (bool, error)
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
