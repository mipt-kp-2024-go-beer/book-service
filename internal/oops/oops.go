package oops

import "github.com/pkg/errors"

// Memory errors
var ErrUnexistedBook = errors.New("Book not found")
var ErrDuplicateID = errors.New("Book with such id already exists")

// Service errors
var ErrLoadBooks = errors.New("Could not load books")
var ErrCreateBook = errors.New("Could not create book")
var ErrUpdateBook = errors.New("Could not update book")
var ErrDeleteBook = errors.New("Could not delete book")

// Real DB specific
var ErrCreatingTable = errors.New("Could not create table")
var ErrDBSetup = errors.New("Could not setup db")

// OS errors
var ErrOSMkdir = errors.New("Could not execute MKDir")
