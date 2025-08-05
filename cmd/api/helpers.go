package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	data envelope,
	headers http.Header,
) error {
	// Encode the data to JSON, returning the error if there was one.
	res, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	res = append(res, '\n')

	// We loop through the header map and add each header to the http.ResponseWriter header map.
	// Note that it's OK if the provided header map is nil.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(res)

	return nil
}

// func (app *application) readJSON(
// 	w http.ResponseWriter,
// 	r *http.Request,
// 	dst any,
// ) error {
// 	return nil
// }