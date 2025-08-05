package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"version": version,
		},
	}

	err := app.writeJSON(w,r,http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w,r,err)
	}
}