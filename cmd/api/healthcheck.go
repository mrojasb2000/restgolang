package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds the information that we want to send in the response.
	data := map[string]string {
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	// Pass the map to the json.Marshal() function. Thist returns a []byte slice
	// containing the encoded JSON. If there was an error, we log it and send the client
	// a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newLine to the JSON. This is just a small nicely to make it easier to
	// view in terminal application
	js = append(js, '\n')

	// At this point we know encoding the data worked without any problems, so we
	// can safely set any necesary HTTP headers for a successful response.
	w.Header().Set("Content-Type", "application/json")

	// Use w.Write() to send the []byte slice containing the JSON as the response body.
	w.Write([]byte(js))
}
