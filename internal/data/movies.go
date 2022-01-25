package data

import "time"

// Annotate the Movie struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID       int64     `json:"id"`         // Unique integer ID for the movie
	CreateAt time.Time `json:"created_at"` // Timestamp for when the movie is added to our database
	Title    string    `json:"title"`      // Movie title
	Year     int32     `json:"year"`       // Movie release year
	Runtime  int32     `json:"runtime"`    // Movie runtime (in minutes)
	Genres   []string  `json:"genres"`     // Slice of genres number for the movie (romance, comedy, etc.)
	Version  int32     `json:"version"`    // time the movie information is updated
}
