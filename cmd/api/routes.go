package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", app.healthCheckHandler)
	mux.HandleFunc("/books/{id}", app.GetBookHandler)
	mux.HandleFunc("/books", app.GetBooksHandler)

	return mux
}