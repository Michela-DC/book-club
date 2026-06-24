package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/Michela-DC/book-club/internal/domain"
)

// CreateBookRequest represents the payload required to create a new book.
// It is typically decoded from the JSON body of an HTTP request.
type CreateBookRequest struct {
	Genre  *string `json:"genre"`
	Year   *int    `json:"year"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Status string  `json:"status"`
}

// UpdateBookRequest represents the payload required to update a book.
type UpdateBookRequest struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Status *string `json:"status"`
	Genre  *string `json:"genre"`
	Year   *int    `json:"year"`
}

// validate checks the fields of CreateBookRequest for correctness.
// It returns an error if any required field is invalid.
func (r *CreateBookRequest) validate() error {
	_, isValidStatus := domain.StringToBookStatusMap[r.Status]

	switch {
	case r.Title == "":
		return errors.New("title cannot be empty")
	case r.Author == "":
		return errors.New("author cannot be empty")
	case r.Genre != nil && *r.Genre == "":
		return errors.New("if specified, genre cannot be empty")
	case !isValidStatus:
		validStatuses := make([]string, 0, len(domain.StringToBookStatusMap))
		for k := range domain.StringToBookStatusMap {
			validStatuses = append(validStatuses, k)
		}
		return fmt.Errorf("status must be one of %v", validStatuses)
	case r.Year != nil && *r.Year > time.Now().Year():
		return errors.New("if specified, year cannot be in the future")
	}

	return nil
}

// validate checks the fields of UpdateBookRequest for correctness.
func (r *UpdateBookRequest) validate() error {
	_, isValidStatus := domain.StringToBookStatusMap[*r.Status]

	switch {
	case r.Title != nil && *r.Title == "":
		return errors.New("if specified, title cannot be empty")
	case r.Author != nil && *r.Author == "":
		return errors.New("if specified, author cannot be empty")
	case r.Genre != nil && *r.Genre == "":
		return errors.New("if specified, genre cannot be empty")
	case !isValidStatus:
		validStatuses := make([]string, 0, len(domain.StringToBookStatusMap))
		for k := range domain.StringToBookStatusMap {
			validStatuses = append(validStatuses, k)
		}
		return fmt.Errorf("status must be one of %v", validStatuses)
	case r.Year != nil && *r.Year > time.Now().Year():
		return errors.New("if specified, year cannot be in the future")
	}

	return nil
}
