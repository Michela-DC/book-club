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
curl -X PUT http://localhost:8080/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "author": "Alan A. A. Donovan & Brian W. Kernighan",
    "title": "The Go Programming Language",
    "year": 2015,
    "status": "SUGGESTED"
  }'

```

Read all Books:
```
curl -X GET http://localhost:8080/v1/books \
  -H "Content-Type: application/json" 

```

Update a Book:
```
curl -X PATCH http://localhost:8080/v1/books/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "author": "Alan A. Donovan & Brian W. Kernighan",
    "title": "The Go Programming Language",
    "year": 2015,
    "status": "SUGGESTED",
    "genre" :  "education"
  }'

```

Delete Book:
```
curl -X DELETE http://localhost:8080/v1/books/{id} \
  -H "Content-Type: application/json" 

```

