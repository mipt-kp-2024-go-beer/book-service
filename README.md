# Book service

## Info

This is a microservice for managing a collection of books.  It provides HTTP operations on books info base.
The Book service can store and tertieve books, with options to filter the list of books using query parameters.

## Features

- Create a new book
- Retrieve all books (filter criteria are optional)
- Retrieve a books using its ID
- Update an existing book
- Delete a book

## Preresquisites

- go version go1.22.9 linux/amd64
- `curl` for testing the API (or any other HTTP client)

## Usage

1. Clone repos:

```bash
usr@usr: git clone https://github.com/mipt-kp-2024-go-beer/book-service
usr@usr: cd book-service
```

2. Synchronize dependencies:

```bash
usr@usr: go mod tidy
```

3. Build and run:

```bash
usr@usr: go build
usr@usr: ./book-service
```

By default, the service will start at `http://127.0.0.1:8080` (you can check it using `netstat` or any other network control application).

## API usage (using `curl`)

### 1. GET /api/books

Retrieve all books. Query parameter `criteria` is optional.

**Example with no `criteria`**

```bash
usr@usr: curl 127.0.0.1:8080/api/v1/books
```

**Example with `criteria`**

```bash
usr@usr: curl 127.0.0.1:8080/api/v1/books?criteria=Alan%20Donovan
```

**Response**:

- Returns a list of books in JSON format
- `null` if there are no any book in database

### 2. GET /api/v1/books/{id}

Retrieve a book by its ID.

**Example with `id = 1`**

```bash
usr@usr: curl 127.0.0.1:8080/api/v1/books/1
```

**Response**

- Returns the book with the specified ID in JSON format.
- Error `Failed to get book by ID...` otherwise

### 3. POST /api/v1/books

Create a new book.

**Example**

```bash
usr@usr: curl 127.0.0.1:8080/api/v1/books/new \
> -H "Content-Type: application/json" \
> -d '{"id": "1", "title": "Go Programming Language", "author": "Alan Donovan", "description": "Good one"}'
```

**Response**

- New book's ID
- Error `Failed to create book...` otherwise

### 4. POST /api/v1/books/{id}

Update an existing book by its ID.

**Example**

```bash
usr@usr: curl -X POST 127.0.0.1:8080/api/v1/books/1 \
> -H "Content-Type: application/json" \
> -d '{"id": "1", "title": "Go Programming Language", "author": "Alan Donovan, Brian Kernighan", "description": "Bad one"}'
```

- Returns the updated book in JSON format
- Error `Failed to update book...` otherwise

### 5. DELETE /api/v1/books/{id}

Delete a book by its ID.

**Example**

```bash
usr@usr: curl -X DELETE 127.0.0.1:8080/api/v1/books/1
```

**Response**
- No response after successful deletion
- Error `Failed to delete book...` otherwise

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
