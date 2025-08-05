package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", app.healthCheckHandler)
	mux.HandleFunc("/books/{id}", app.GetBookHandler)

	return mux
}