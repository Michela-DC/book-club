package controller

// CreateBookRequest represents the payload required to create a new book.
// It is typically decoded from the JSON body of an HTTP request.
type CreateBookRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Year   int    `json:"year"`
}

// validate checks the fields of CreateBookRequest for correctness.
// It returns an error if any required field is invalid.
// Currently, it is not implemented.
func (*CreateBookRequest) validate() error {
	// TODO: IMPLEMENT
	return nil
}
