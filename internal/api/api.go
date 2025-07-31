package api

import (
	"encoding/json"
	"net/http"

	"github.com/daalfox/bookstore/internal/spec"
)

type BookstoreApi struct{}

func NewBookstoreApi() BookstoreApi {
	return BookstoreApi{}
}

func (h *BookstoreApi) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := []spec.Book{}
	json.NewEncoder(w).Encode(books)
}
