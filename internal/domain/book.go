package domain

type Book struct {
	ID            string     `json:"id"`
	Title         string     `json:"author"`
	Author        string     `json:"title"`
	PublishedYear int        `json:"year"`
	Status        BookStatus `json:"status"`
	//TODO: ADD GENRE
}

type BookStatus string

const (
	BookStatusSuggested BookStatus = "SUGGESTED"
	BookStatusReading   BookStatus = "READING"
	BookStatusDiscarded BookStatus = "DISCARDED"
	BookStatusCompleted BookStatus = "COMPLETED"
	BookStatusSaved     BookStatus = "SAVED"
)

var StringToBookStatusMap = map[string]BookStatus{
	"SUGGESTED": BookStatusSuggested,
	"READING":   BookStatusReading,
	"DISCARDED": BookStatusDiscarded,
	"COMPLETED": BookStatusCompleted,
	"SAVED":     BookStatusSaved,
}

type BookFilters struct {
	ID            *string
	Title         *string
	Author        *string
	PublishedYear *int
	Status        *BookStatus
}

type BookRepository interface {
	Create(*Book) (*Book, error)
	List(*BookFilters) ([]*Book, error)
	Update(*Book) error
	Delete(string) error
}
