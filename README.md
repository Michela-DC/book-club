# book-club

A simple book club REST API following the Clean Architecture principles.

The API provides CRUD operations to manage a book club reading suggestions.

## Run the API

Prerequisites:
- go
- sqlite3

```
go run ./cmd/book-club
```

## Sample Requests

Create a Book:
```
curl -X PUT http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "author": "Alan A. A. Donovan & Brian W. Kernighan",
    "title": "The Go Programming Language",
    "year": 2015,
    "status": "SUGGESTED"
  }'

```


