package interactor

import (
	"context"
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
func (b *BookInteractor) CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book == nil {
		return nil, errors.New("empty book info")
	}
	if book.Status == domain.BookStatusCompleted || book.Status == domain.BookStatusDiscarded {
		return nil, errors.New("cannot create book with status " + string(book.Status))
	}
	return b.repo.Create(ctx, book)
}

// ReadBooks retrieves a list of books that match the provided filters.
// TODO: servirá modificare la firma dei metodi con (b *BookInteractor).
func (*BookInteractor) ReadBooks(_ *domain.BookFilters) ([]*domain.Book, error) {
	return nil, errors.New("not implemented")
}

// UpdateBook updates the information of an existing book in the repository.
// TODO: servirá modificare la firma dei metodi con (b *BookInteractor).
func (b *BookInteractor) UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	if book == nil {
		return nil, errors.New("book not found")
	}

	return book, b.repo.Update(ctx, book)
}

// DeleteBook removes a book from the repository by its unique ID.
// TODO: servirá modificare la firma dei metodi con (b *BookInteractor).
func (*BookInteractor) DeleteBook(_ string) error {
	return errors.New("not implemented")
}
