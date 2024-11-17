package library

import (
	"context"
	"fmt"
)

type AppStockService struct {
	store StockStore
}

func NewStockService(s StockStore) *AppStockService {
	return &AppStockService{store: s}
}

func (s *AppStockService) GetStock(ctx context.Context, bookID string) (*Stock, error) {
	// Fetch stock from the store
	stock, err := s.store.LoadStock(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("could not load stock for book with id %s: %w", bookID, err)
	}
	return stock, nil
}

func (s *AppStockService) ChangeStock(ctx context.Context, bookID string, delta int) error {
	// Fetch the current stock record from the store
	stock, err := s.store.LoadStock(ctx, bookID)
	if err != nil {
		return fmt.Errorf("could not load stock for book with id %s: %w", bookID, err)
	}

	// Check for overflow
	if stock.AvailableStock+delta < 0 {
		return fmt.Errorf("unavailable delta %d for book with id %s", delta, bookID)
	}

	// It looks like everything is ok. Let's try to update the stock
	if err := s.store.UpdateStock(ctx, bookID, delta); err != nil {
		return fmt.Errorf("could not update stock for book with id %s: %w", bookID, err)
	}

	return nil
}
