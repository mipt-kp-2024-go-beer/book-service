package library

import "context"

// Stock represents the stock record of a book
type Stock struct {
	BookID         string
	TotalStock     int
	LentStock      int
	AvailableStock int
}

// StockService defines the interface for interacting with stock (business logic)
type StockService interface {
	GetStock(ctx context.Context, bookID string) (*Stock, error)
	SaveStock(ctx context.Context, stock Stock) error
	ChangeStock(ctx context.Context, bookID string, delta int) error
}

// StockStore defines the interface for database interactions related to stock
type StockStore interface {
	LoadStock(ctx context.Context, bookID string) (*Stock, error)
	SaveStock(ctx context.Context, stock Stock) error
	UpdateStock(ctx context.Context, bookID string, delta int) error
}
