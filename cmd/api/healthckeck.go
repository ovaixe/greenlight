package main

import (
	"net/http"
)

// Declare a handler whcih writes a plain-text response with information about the application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
