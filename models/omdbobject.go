package models

import "fmt"

// OMDBObject struct contains all the information that is found on both a movie and TV show JSON obtained from the OMDB API
type OMDBObject struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Awards   string
	Ratings  Ratings
}

// PrintOMDBInfo takes the information stored in an OMDBObject struct and prints it out to the user
func (o OMDBObject) PrintOMDBInfo() {
	fmt.Println("Title:", o.Title)
	fmt.Println("Year:", o.Year)
	fmt.Println("Rated:", o.Rated)
	fmt.Println("Released:", o.Released)
	fmt.Println("Runtime:", o.Runtime)
	fmt.Println("Genre:", o.Genre)
	fmt.Println("Director:", o.Director)
	fmt.Println("Writer:", o.Writer)
	fmt.Println("Cast:", o.Actors)
	fmt.Println("Plot:", o.Plot)
	fmt.Println("Language:", o.Language)
	fmt.Println("Awards:", o.Awards)
}
