package controller

// CreateBookRequest represents the payload required to create a new book.
// It is typically decoded from the JSON body of an HTTP request.
type CreateBookRequest struct {
	Author string `json:"author"` // Author of the book
	Title  string `json:"title"`  // Title of the book
	Year   int    `json:"year"`   // Publication year of the book
	Status string `json:"status"` // Reading status of the book (e.g., "reading", "completed")
}

// validate checks the fields of CreateBookRequest for correctness.
// It returns an error if any required field is invalid.
// Currently, it is not implemented.
func (b *CreateBookRequest) validate() error {
	// TODO: IMPLEMENT
	return nil
}
