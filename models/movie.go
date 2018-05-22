package models

import "fmt"

// Movie struct contains all the information of a movie obtained
// from the OMDB API
type Movie struct {
	Title string
	/*Year       int
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Cast       string
	Plot       string
	Language   string
	Awards     string
	Ratings    []Rating
	Production string*/
}

// GetMovie performs an HTTP request to the OMDB API to get a movie's information and then print it to the user
func GetMovie(name string) {
	fmt.Println(name, "test")
}
