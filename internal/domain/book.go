package domain

import "context"

// Book is the main entity.
type Book struct {
	ID            string     `json:"id"`
	Title         string     `json:"author"`
	Author        string     `json:"title"`
	Genre         *string    `json:"genre"`
	PublishedYear *int       `json:"year"`
	Status        BookStatus `json:"status"`
}

// BookStatus defines the current book status for the book club.
// Book lifecycle: SAVED (optional status), SUGGESTED, DISCARDED or READING, COMPLETED.
type BookStatus string

const (
	// BookStatusSuggested reppresents the book is beeing considered for reading.
	BookStatusSuggested BookStatus = "SUGGESTED"
	// BookStatusReading means the book has been accepted for reading and is currently being read.
	BookStatusReading BookStatus = "READING"
	// BookStatusDiscarded means the book has not been approved for reading.
	BookStatusDiscarded BookStatus = "DISCARDED"
	// BookStatusCompleted means the book has been read.
	BookStatusCompleted BookStatus = "COMPLETED"
	// BookStatusSaved reppresents books that are possible new candidates for reading.
	BookStatusSaved BookStatus = "SAVED"
)

// StringToBookStatusMap maps a valid string reppresentation of a book status to the correct [BookStatus] variable.
var StringToBookStatusMap = map[string]BookStatus{
	"SUGGESTED": BookStatusSuggested,
	"READING":   BookStatusReading,
	"DISCARDED": BookStatusDiscarded,
	"COMPLETED": BookStatusCompleted,
	"SAVED":     BookStatusSaved,
}

// BookFilters represents the possible filters to use for searching.
type BookFilters struct {
	ID            *string
	Title         *string
	Author        *string
	Genre         *string
	PublishedYear *int
	Status        *BookStatus
}

// BookRepository is the book persistency repository.
// BookRepository defines the interface for persisting and retrieving books.
// It abstracts the underlying storage mechanism (e.g., SQLite, in-memory, etc.).
type BookRepository interface {
	// Create inserts a new book into the repository.
	Create(ctx context.Context, book *Book) (*Book, error)
	// List retrieves all books matching the provided filters.
	List(ctx context.Context, filters *BookFilters) ([]*Book, error)
	// Update modifies an existing book in the repository.
	Update(ctx context.Context, book *Book) error
	// Delete removes a book identified by its unique ID from the repository.
	Delete(ctx context.Context, id string) error
}
