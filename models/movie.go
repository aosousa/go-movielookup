package models

import "fmt"

// Movie struct contains all the information of a movie obtained
// from the OMDB API
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Awards     string
	Ratings    []Rating
	Production string
}

// PrintMovie takes the information stored in a Movie struct and prints it out to the user
func (m Movie) PrintMovie() {
	fmt.Println("Title:", m.Title)
	fmt.Println("Year:", m.Year)
	fmt.Println("Rated:", m.Rated)
	fmt.Println("Released:", m.Released)
	fmt.Println("Runtime:", m.Runtime)
	fmt.Println("Genre:", m.Genre)
	fmt.Println("Director:", m.Director)
	fmt.Println("Writer:", m.Writer)
	fmt.Println("Cast:", m.Actors)
	fmt.Println("Plot:", m.Plot)
	fmt.Println("Language:", m.Language)
	fmt.Println("Awards:", m.Awards)
	fmt.Println("Production:", m.Production)

	if len(m.Ratings) > 0 {
		fmt.Println("Ratings:")
		for i := range m.Ratings {
			rating := "* " + m.Ratings[i].Source + ": " + m.Ratings[i].Value
			fmt.Println(rating)
		}
	} else {
		fmt.Println("Ratings: N/A")
	}
}
