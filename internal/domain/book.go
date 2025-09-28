package domain

// Book is the main entity.
type Book struct {
	ID            string     `json:"id"`
	Title         string     `json:"author"`
	Author        string     `json:"title"`
	PublishedYear int        `json:"year"`
	Status        BookStatus `json:"status"`
	//TODO: ADD GENRE
}

// BookStatus defines the current book status for the book club. 
// Book lifecycle: SAVED (optional status), SUGGESTED, DISCARDED or READING, COMPLETED.
type BookStatus string

const (
	// BookStatusSuggested reppresents the book is beeing considered for reading.
	BookStatusSuggested BookStatus = "SUGGESTED"
	// BookStatusReading means the book has been accepted for reading and is currently being read.
	BookStatusReading   BookStatus = "READING"
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
	PublishedYear *int
	Status        *BookStatus
}


// BookRepository is the book persistency repository
type BookRepository interface {
	Create(*Book) (*Book, error)
	List(*BookFilters) ([]*Book, error)
	Update(*Book) error
	Delete(string) error
}
