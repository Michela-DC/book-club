package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/Michela-DC/book-club/internal/domain"
)

// BookInteractor defines the application logic for managing books.
type BookInteractor interface {
	// CreateBook creates a new book and persists it in the data store.
	CreateBook(book *domain.Book) (*domain.Book, error)
	// ReadBooks retrieves a list of books that match the provided filters.
	ReadBooks(filters *domain.BookFilters) ([]*domain.Book, error)
	// UpdateBook updates the information of an existing book in the repository.
	UpdateBook(book *domain.Book) (*domain.Book, error)
	// DeleteBook removes a book from the repository by its unique ID.
	DeleteBook(id string) error
}

// BookController implements [webservice.CRUDController] to handle
// HTTP requests related to book resources.
type BookController struct {
	interactor BookInteractor
	logger     *slog.Logger
}

// NewBookController creates a new BookController with the given interactor and logger.
func NewBookController(i BookInteractor, l *slog.Logger) *BookController {
	return &BookController{
		interactor: i,
		logger:     l,
	}
}

// Create handles HTTP requests for creating a new book. It decodes
// the request body, validates the input, creates the book via the
// interactor, and writes the created book as JSON to the response.
func (b *BookController) Create(w http.ResponseWriter, r *http.Request) {
	var cbr CreateBookRequest
	err := json.NewDecoder(r.Body).Decode(&cbr)
	if err != nil {
		b.logger.With("error", err).Error("unable to get request body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = cbr.validate()
	if err != nil {
		b.logger.With("error", err).Error("invalid request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, ok := domain.StringToBookStatusMap[cbr.Status]
	if !ok {
		b.logger.With("status", cbr.Status).Error("invalid book status")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book, err := b.interactor.CreateBook(&domain.Book{
		ID:            uuid.NewString(),
		Title:         cbr.Title,
		Author:        cbr.Author,
		PublishedYear: cbr.Year,
		Status:        status,
	})
	if err != nil {
		b.logger.With("error", err).Error("unable to create book")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		b.logger.With("error", err).Error("unable to encode book")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// Read handles HTTP requests for retrieving books.
// Currently, it is not implemented and always returns 501 Not Implemented.
func (*BookController) Read(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// Update handles HTTP requests for updating an existing book.
// Currently, it is not implemented and always returns 501 Not Implemented.
func (*BookController) Update(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// Delete handles HTTP requests for deleting a book by ID.
// Currently, it is not implemented and always returns 501 Not Implemented.
func (*BookController) Delete(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
