package interactor

import (
	"errors"
	"log/slog"

	"github.com/Michela-DC/book-club/internal/domain"
)

type BookInteractor struct {
	repo   domain.BookRepository
	logger *slog.Logger
}

func NewBookInteractor(repo domain.BookRepository, logger *slog.Logger) *BookInteractor {
	return &BookInteractor{
		repo:   repo,
		logger: logger,
	}
}

// define CreateBook as a function of BookInteractor, b is a pointer receiver
func (b *BookInteractor) CreateBook(book *domain.Book) (*domain.Book, error) {
	if book == nil {
		return nil, errors.New("empty book info")
	}
	if book.Status == domain.BookStatusCompleted || book.Status == domain.BookStatusDiscarded {
		return nil, errors.New("cannot create book with status " + string(book.Status))
	}
	return b.repo.Create(book)
}
