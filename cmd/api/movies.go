package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mrojasb2000/greenlight/internal/data"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder respose.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
} 

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// We can then use the ByNames() method to get the value of the "id" parameter from
	// the slice. In our project all movies will have a unique positive integer ID, but
	// the value returned by ByName() is always a string. So we try to convert it to a
	// base 10 integer (with a bit size of 64). If the parameter couldn't be converted,
	// or is less than 1, we know the ID is invalid so we use the http.NotFound()
	// function to return a 404 Not Found response.
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID we extrated from
	// the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.
	movie := data.Movie {
		ID: id,
		CreateAt: time.Now(),
		Title: "Casablanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}

	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encontrered a problem and could not process your request", http.StatusInternalServerError)
	}
}