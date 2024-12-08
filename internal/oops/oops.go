package oops

import "github.com/pkg/errors"

// Memory errors
var ErrUnexistedBook = errors.New("Book not found")
var ErrDuplicateID = errors.New("Book with such id already exists")

// Service errors
var ErrLoadBooks = errors.New("Could not load books")
var ErrCreateBooks = errors.New("Could not create book")
