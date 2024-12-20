package memory

import (
	"context"
	"sync"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/oops"
)

type MemoryBookStore struct {
	mu    sync.RWMutex
	books map[string]library.Book
}

func NewMemoryBookStore() *MemoryBookStore {
	return &MemoryBookStore{
		books: make(map[string]library.Book),
	}
}

func (s *MemoryBookStore) LoadBooks(ctx context.Context, criteria string) ([]library.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []library.Book
	for _, book := range s.books {
		// Lookup for the same substring in in book title, author or decription
		if strContains(book.Title, criteria) || strContains(book.Author, criteria) || strContains(book.Description, criteria) {
			result = append(result, book)
		}
	}
	return result, nil
}

func (s *MemoryBookStore) LoadBookByID(ctx context.Context, id string) (*library.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, exists := s.books[id]
	if !exists {
		return nil, oops.ErrUnexistedBook
	}
	return &book, nil
}
func (s *MemoryBookStore) SaveBook(ctx context.Context, book library.Book) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[book.ID]; exists {
		return "", oops.ErrDuplicateID
	}

	s.books[book.ID] = book
	return book.ID, nil
}

func (s *MemoryBookStore) UpdateBook(ctx context.Context, id string, book library.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return oops.ErrUnexistedBook
	}

	s.books[id] = book
	return nil
}

func (s *MemoryBookStore) DeleteBook(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return oops.ErrUnexistedBook
	}

	delete(s.books, id)
	return nil
}

// Simple search in db
func strContains(str, substr string) bool {
	if substr == "" {
		return true
	}

	if str == "" {
		return false
	}

	// Substr length should be less or equal to str length
	if len(str) < len(substr) {
		return false
	}

	// Lookup for substring
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}
