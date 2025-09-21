package webservice

import (
	"net/http"
)

type BookController interface {
	Create(http.ResponseWriter, *http.Request)
}

func NewHandler(bookController BookController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /books", bookController.Create)

	return mux
}
