package interactor

import (
	"errors"
	"log/slog"

	"github.com/Michela-DC/book-club/internal/domain"
)

// BookInteractor provides the application logic for managing books.
// It coordinates between the domain layer and repositories.
type BookInteractor struct {
	repo   domain.BookRepository
	logger *slog.Logger
}

// NewBookInteractor creates a new BookInteractor with the given repository and logger.
func NewBookInteractor(repo domain.BookRepository, logger *slog.Logger) *BookInteractor {
	return &BookInteractor{
		repo:   repo,
		logger: logger,
	}
}

// CreateBook validates the provided book and delegates its creation
// to the underlying repository. It returns an error if the book is nil,
// or if the status is invalid for creation (e.g., "completed" or "discarded").
func (b *BookInteractor) CreateBook(book *domain.Book) (*domain.Book, error) {
	if book == nil {
		return nil, errors.New("empty book info")
	}
	if book.Status == domain.BookStatusCompleted || book.Status == domain.BookStatusDiscarded {
		return nil, errors.New("cannot create book with status " + string(book.Status))
	}
	return b.repo.Create(book)
}

func (b *BookInteractor) ReadBooks(filters *domain.BookFilters) ([]*domain.Book, error) {
	return nil, errors.New("not implemented")
}
	
func (b *BookInteractor) UpdateBook(book *domain.Book) (*domain.Book, error) {
	return nil, errors.New("not implemented")
}
	
func (b *BookInteractor) DeleteBook(id string) (error) {
	return errors.New("not implemented")
}
