package library_test

import (
	"context"
	"testing"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/library/memory"
)

func TestBookService(t *testing.T) {
	bookStore := memory.NewMemoryBookStore()
	bookService := library.NewBookService(bookStore)

	// Test 1: Add a new book
	t.Run("CreateBook", func(t *testing.T) {
		book := library.Book{
			ID:          "1",
			Title:       "Go Programming",
			Author:      "John Doe",
			Description: "A book about Go programming",
		}

		id, err := bookService.CreateBook(context.Background(), book)
		if err != nil {
			t.Errorf("Failed create book with id %s: %s", book.ID, err)
		}
		if id != book.ID {
			t.Errorf("CreateBook failed, expected id %s: got %s", book.ID, id)
		}

		savedBook, err := bookStore.LoadBookByID(context.Background(), book.ID)
		if err != nil {
			t.Errorf("LoadBookByID failed: %s", err)
		}
		if book != *savedBook {
			t.Errorf("LoadBookByID failed: expected %+v but got %+v", book, *savedBook)
		}
	})

	// Test 2: Get a book by ID
	t.Run("GetBookByID", func(t *testing.T) {
		book := library.Book{
			ID:          "2",
			Title:       "Advanced Go",
			Author:      "Jane Smith",
			Description: "An advanced Go programming guide",
		}

		_, err := bookService.CreateBook(context.Background(), book)
		if err != nil {
			t.Errorf("Failed create book with id %s: %s", book.ID, err)
		}

		fetchedBook, err := bookService.GetBookByID(context.Background(), "2")
		if err != nil {
			t.Errorf("Failed to get book with id %s: %s", "2", err)
		}

		if book != *fetchedBook {
			t.Errorf("Failed to get book by id %s after creation", book.ID)
		}
	})

	// Test 3: Get books by search criteria
	t.Run("GetBooks", func(t *testing.T) {
		book1 := library.Book{
			ID:          "3",
			Title:       "Go for Beginners",
			Author:      "Alice Brown",
			Description: "A beginner's guide to Go",
		}
		book2 := library.Book{
			ID:     "4",
			Title:  "The C++ Programming Language (4th Edition)",
			Author: "Bjarne Stroustrup",
			Description: `The C++ Programming Language, Fourth Edition, delivers meticulous,
						  "richly explained, and integrated coverage of the entire language—its facilities,
						  abstraction mechanisms, standard libraries, and key design techniques.`,
		}
		_, err := bookService.CreateBook(context.Background(), book1)
		if err != nil {
			t.Errorf("Failed create book with id %s: %s", book1.ID, err)
		}

		_, err = bookService.CreateBook(context.Background(), book2)
		if err != nil {
			t.Errorf("Failed create book with id %s: %s", book2.ID, err)
		}

		books, err := bookService.GetBooks(context.Background(), "C++")
		if err != nil {
			t.Errorf("Failed to find out the book with criteria %s: %s", "C++", err)
		}
		if len(books) != 1 {
			t.Errorf("Wrong GetBooks answer")
		}

		books, err = bookService.GetBooks(context.Background(), "Advanced")
		if err != nil {
			t.Errorf("Failed to find out the book with criteria %s: %s", "Advanced", err)
		}
		if len(books) != 1 {
			t.Errorf("Wrong GetBooks answer")
		}
	})

	t.Run("UpdateBook", func(t *testing.T) {
		book := library.Book{
			ID:          "5",
			Title:       "Intro to Go",
			Author:      "Chris White",
			Description: "A basic introduction to Go",
		}

		_, err := bookService.CreateBook(context.Background(), book)
		if err != nil {
			t.Errorf("Couldn't create book with id %s: %s", book.ID, err)
		}

		updatedBook := book
		updatedBook.Title = "Introduction to Go"
		err = bookService.UpdateBook(context.Background(), book.ID, updatedBook)
		if err != nil {
			t.Errorf("Couldn't update book info with id %s: %s", book.ID, err)
		}

		updatedBookFromStore, err := bookStore.LoadBookByID(context.Background(), book.ID)
		if err != nil {
			t.Errorf("Couldn't find the book with id %s after update: %s", book.ID, err)
		}

		if updatedBook != *updatedBookFromStore {
			t.Errorf("Failed update the book")
		}
	})

	t.Run("DeleteBook", func(t *testing.T) {
		book := library.Book{
			ID:          "6",
			Title:       "Learn Go",
			Author:      "David Black",
			Description: "A comprehensive Go programming guide",
		}

		_, err := bookService.CreateBook(context.Background(), book)
		if err != nil {
			t.Errorf("Couldn't create book with id %s: %s", book.ID, err)
		}

		err = bookService.DeleteBook(context.Background(), book.ID)
		if err != nil {
			t.Errorf("Couldn't delete book with id %s: %s", book.ID, err)
		}

		_, err = bookStore.LoadBookByID(context.Background(), book.ID)
		if err == nil {
			t.Errorf("Book with id %s was not deleted", book.ID)
		}
	})
}
