package library_test

import (
	"context"
	"testing"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/library/memory"
)

func TestStockService(t *testing.T) {
	stockStore := memory.NewMemoryStockStore()
	stockService := library.NewStockService(stockStore)

	// Test 1: Save and get stock for a book
	t.Run("SaveAndGetStock", func(t *testing.T) {
		stock := library.Stock{
			BookID:         "1",
			AvailableStock: 10,
			LentStock:      2,
		}

		err := stockService.SaveStock(context.Background(), stock)
		if err != nil {
			t.Errorf("SaveStock() failed for stock %+v with error %s", stock, err)
		}

		storedStock, err := stockService.GetStock(context.Background(), stock.BookID)
		if err != nil {
			t.Errorf("GetStock() failed for stock %+v with error %s", stock, err)
		}

		if stock != *storedStock {
			t.Errorf("Wrong answer from GetStock(): expected %+v, got %+v", stock, *storedStock)
		}
	})

	// Test 2: Change stock (increase available stock)
	t.Run("ChangeStock_Increase", func(t *testing.T) {
		stock := library.Stock{
			BookID:         "2",
			AvailableStock: 5,
			LentStock:      0,
		}

		err := stockService.SaveStock(context.Background(), stock)
		if err != nil {
			t.Errorf("SaveStock() failed for stock %+v with error %s", stock, err)
		}

		err = stockService.ChangeStock(context.Background(), stock.BookID, 3)
		if err != nil {
			t.Errorf("ChangeStock() failed for stock %+v with error %s", stock, err)
		}

		storedStock, err := stockService.GetStock(context.Background(), stock.BookID)
		if err != nil {
			t.Errorf("GetStock() failed for stock %+v with error %s", stock, err)
		}

		if storedStock.AvailableStock != 8 {
			t.Errorf("Wrong answer after ChangeStock(): expected %d, got %d\n", 8, stock.AvailableStock)
		}
	})

	// Test 3: Change stock (decrease available stock)
	t.Run("ChangeStock_Decrease", func(t *testing.T) {
		stock := library.Stock{
			BookID:         "3",
			AvailableStock: 5,
			LentStock:      0,
		}

		err := stockService.SaveStock(context.Background(), stock)
		if err != nil {
			t.Errorf("SaveStock() failed for stock %+v with error %s", stock, err)
		}

		err = stockService.ChangeStock(context.Background(), stock.BookID, -2)
		if err != nil {
			t.Errorf("ChangeStock() failed for stock %+v with error %s", stock, err)
		}

		storedStock, err := stockService.GetStock(context.Background(), stock.BookID)
		if err != nil {
			t.Errorf("GetStock() failed for stock %+v with error %s", stock, err)
		}

		if storedStock.AvailableStock != 3 {
			t.Errorf("Wrong answer after ChangeStock(): expected %d, got %d\n", 3, stock.AvailableStock)
		}
	})

	// Test 4: Change stock (not enough available stock)
	t.Run("ChangeStock_InsufficientStock", func(t *testing.T) {
		stock := library.Stock{
			BookID:         "4",
			AvailableStock: 3,
			LentStock:      0,
		}

		err := stockService.SaveStock(context.Background(), stock)
		if err != nil {
			t.Errorf("SaveStock() failed for stock %+v with error %s", stock, err)
		}

		// Attempt to overflow the field
		err = stockService.ChangeStock(context.Background(), stock.BookID, -5)
		if err == nil {
			t.Errorf("ChangeStock() failed for stock %+v with error %s", stock, err)
		}
	})
}
