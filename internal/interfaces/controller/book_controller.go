package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/Michela-DC/book-club/internal/domain"
	"github.com/Michela-DC/book-club/internal/infrastructure/db"
	"github.com/Michela-DC/book-club/pkg/utilities"
)

// BookInteractor defines the application logic for managing books.
type BookInteractor interface {
	// CreateBook creates a new book and persists it in the data store.
	CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// ReadBooks retrieves a Read of books that match the provided filters.
	ReadBooks(filters *domain.BookFilters) ([]*domain.Book, error)
	// UpdateBook updates the information of an existing book in the repository.
	UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	// DeleteBook removes a book from the repository by its unique ID.
	DeleteBook(ctx context.Context, id string) error
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
// interactor and writes the created book as JSON to the response.
func (b *BookController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	status := domain.StringToBookStatusMap[cbr.Status]

	book, err := b.interactor.CreateBook(ctx, &domain.Book{
		ID:            uuid.NewString(),
		Title:         cbr.Title,
		Author:        cbr.Author,
		Genre:         cbr.Genre,
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
func (b *BookController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var ubr UpdateBookRequest
	err := json.NewDecoder(r.Body).Decode(&ubr)
	if err != nil {
		b.logger.With("error", err).Error("unable to get request body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ubr.validate()
	if err != nil {
		b.logger.With("error", err).Error("invalid request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var status domain.BookStatus
	if ubr.Status != nil {
		var ok bool
		status, ok = domain.StringToBookStatusMap[*ubr.Status]
		if !ok {
			b.logger.With("status", ubr.Status).Error("invalid book status")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	book, err := b.interactor.UpdateBook(ctx, &domain.Book{
		ID:            uuid.NewString(),
		Title:         utilities.Optional(ubr.Title),
		Author:        utilities.Optional(ubr.Author),
		Genre:         ubr.Genre,
		PublishedYear: ubr.Year,
		Status:        status,
	})
	if err != nil {
		b.logger.With("error", err).Error("unable to update book")
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

// Delete handles HTTP requests for deleting a book by ID.
func (b *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: handle not found
	ctx := r.Context()
	bookID := r.PathValue("id")

	err := b.interactor.DeleteBook(ctx, bookID)

	if err != nil {
		b.logger.With("error", err).Error("unable to delete book")
		if errors.Is(err, db.ErrorNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b.logger.With("id", bookID).Info("book deleted")
}
