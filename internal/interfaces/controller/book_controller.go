package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/Michela-DC/book-club/internal/domain"
)

type BookInteractor interface {
	CreateBook(book *domain.Book) (*domain.Book, error)
}

// BookController implements [webservice.CRUDController].
type BookController struct {
	interactor BookInteractor
	logger     *slog.Logger
}

func NewBookController(i BookInteractor, l *slog.Logger) *BookController {
	return &BookController{
		interactor: i,
		logger:     l,
	}
}

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

func (b *BookController) Read(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (b *BookController) Update(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (b *BookController) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
