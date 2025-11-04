package webservice

import (
	"net/http"
)

// CRUDController defines the basic Create, Read, Update, and Delete
// operations for handling HTTP requests in a RESTful API.
type CRUDController interface {
	// Create handles the HTTP request to create a new resource.
	Create(w http.ResponseWriter, r *http.Request)
	// Read handles the HTTP request to retrieve one or more resources.
	Read(w http.ResponseWriter, r *http.Request)
	// Update handles the HTTP request to modify an existing resource.
	Update(w http.ResponseWriter, r *http.Request)
	// Delete handles the HTTP request to remove an existing resource.
	Delete(w http.ResponseWriter, r *http.Request)
}

// NewHandler registers the CRUDController routes for book resources and
// returns an http.Handler. It maps each HTTP method and endpoint to the
// corresponding CRUD operation.
func NewHandler(bookController CRUDController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /v1/books", bookController.Create)
	mux.HandleFunc("GET /v1/books", bookController.Read)
	mux.HandleFunc("PATCH /v1/books/{id}", bookController.Update)
	mux.HandleFunc("DELETE /v1/books/{id}", bookController.Delete)
	return mux
}
