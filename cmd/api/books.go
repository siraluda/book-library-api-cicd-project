package main

import (
	"net/http"
	"strconv"
)

// Get a specific book with the GET /v1/books/:id endpoint
func (app application) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	book, err := app.models.Books.Get(bookId)
	if err != nil {
		app.notFoundResponse(w,r)
		return
	}

	err = app.writeJSON(w,r,http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app application) GetBooksHandler(w http.ResponseWriter, r *http.Request)  {
	books, err := app.models.Books.GetAll()
	if err != nil {
		app.notFoundResponse(w,r)
		return
	}

	err = app.writeJSON(w,r,http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w,r,err)
		return
	}
}
