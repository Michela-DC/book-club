package webservice

import (
	"net/http"
)

type CRUDController interface {
	Create(http.ResponseWriter, *http.Request)
	Read(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func NvewHandler(bookController CRUDController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /v1/books", bookController.Create)
	mux.HandleFunc("GET /v1/books", bookController.Read)
	mux.HandleFunc("PATCH /v1/books/{id}",bookController.Update)
	mux.HandleFunc("DELETE /v1/books/{id}", bookController.Delete)

	return mux
}
