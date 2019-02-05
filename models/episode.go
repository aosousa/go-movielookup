package models

import "fmt"

// Episode struct contains all the information of a TV show episode obtained from the OMDB API
type Episode struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Season   string
	Number   string `json:"Episode"` // need to set this json value because it differs from the property name
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Plot     string
	Language string
	Ratings  Ratings
}

// PrintEpisode takes the information stored in an Episode struct and prints it out to the user
func (e Episode) PrintEpisode() {
	fmt.Println("Title:", e.Title)
	fmt.Println("Year:", e.Year)
	fmt.Println("Rated:", e.Rated)
	fmt.Println("Released:", e.Released)
	fmt.Println("Runtime:", e.Runtime)
	fmt.Println("Genre:", e.Genre)
	fmt.Println("Director:", e.Director)
	fmt.Println("Writer:", e.Writer)
	fmt.Println("Plot:", e.Plot)
	fmt.Println("Language:", e.Language)

	e.Ratings.PrintRatings()
}
