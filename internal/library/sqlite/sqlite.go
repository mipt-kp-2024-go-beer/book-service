package sqlite

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mipt-kp-2024-go-beer/book-service/internal/library"
	"github.com/mipt-kp-2024-go-beer/book-service/internal/oops"
	"github.com/pkg/errors"
)

type SQLiteBookStore struct {
	db *sql.DB
}

func NewSQLiteBookStore(path string) (*SQLiteBookStore, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, oops.ErrOSMkdir.Error())
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create the books table if it doesn't exist
	create := `CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT,
		description TEXT
	);`

	_, err = db.Exec(create)
	if err != nil {
		return nil, errors.Wrap(err, oops.ErrCreatingTable.Error())
	}

	return &SQLiteBookStore{db: db}, nil
}

func (s *SQLiteBookStore) LoadBooks(ctx context.Context, criteria string) ([]library.Book, error) {
	query := `SELECT id, title, author, description FROM books WHERE title LIKE ? OR author LIKE ? OR description LIKE ?`
	rows, err := s.db.QueryContext(ctx, query, "%"+criteria+"%", "%"+criteria+"%", "%"+criteria+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []library.Book
	for rows.Next() {
		var book library.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (s *SQLiteBookStore) LoadBookByID(ctx context.Context, id string) (*library.Book, error) {
	query := `SELECT id, title, author, description FROM books WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, id)

	var book library.Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, oops.ErrUnexistedBook
		}
		return nil, err
	}

	return &book, nil
}

func (s *SQLiteBookStore) SaveBook(ctx context.Context, book library.Book) (string, error) {
	query := `INSERT INTO books (id, title, author, description) VALUES (?, ?, ?, ?)`
	_, err := s.db.ExecContext(ctx, query, book.ID, book.Title, book.Author, book.Description)
	if err != nil {
		return "", err
	}

	return book.ID, nil
}

func (s *SQLiteBookStore) UpdateBook(ctx context.Context, id string, book library.Book) error {
	query := `UPDATE books SET title = ?, author = ?, description = ? WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, book.Title, book.Author, book.Description, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return oops.ErrUnexistedBook
	}

	return nil
}

func (s *SQLiteBookStore) DeleteBook(ctx context.Context, id string) error {
	query := `DELETE FROM books WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return oops.ErrUnexistedBook
	}

	return nil
}
