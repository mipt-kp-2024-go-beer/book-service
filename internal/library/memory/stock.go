package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
)

// InMemoryStockStore is an in-memory implementation of the StockStore interface
type MemoryStockStore struct {
	mu     sync.RWMutex
	stocks map[string]library.Stock // Maps book ID to stock info
}

// NewInMemoryStockStore creates a new instance of InMemoryStockStore
func NewMemoryStockStore() *MemoryStockStore {
	return &MemoryStockStore{
		stocks: make(map[string]library.Stock),
	}
}

// LoadStock retrieves the stock information for a specific book from the in-memory store
func (s *MemoryStockStore) LoadStock(ctx context.Context, bookID string) (*library.Stock, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stock, exists := s.stocks[bookID]
	if !exists {
		return nil, fmt.Errorf("stock for book with id %s not found", bookID)
	}
	return &stock, nil
}

// SaveStock creates a new stock record for a specific book in the in-memory store
func (s *MemoryStockStore) SaveStock(ctx context.Context, stock library.Stock) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.stocks[stock.BookID]; exists {
		return fmt.Errorf("stock for book with id %s already exists", stock.BookID)
	}

	s.stocks[stock.BookID] = stock
	return nil
}

// UpdateStock adjusts the stock record for a specific book by a given delta
func (s *MemoryStockStore) UpdateStock(ctx context.Context, bookID string, delta int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stock, exists := s.stocks[bookID]
	if !exists {
		return fmt.Errorf("stock for book with id %s not found", bookID)
	}

	stock.AvailableStock += delta
	if stock.AvailableStock < 0 {
		return fmt.Errorf("not enough stock for book with id %s", bookID)
	}

	s.stocks[bookID] = stock
	return nil
}
