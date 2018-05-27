package models

import "fmt"

// Show struct contains all the information of a TV show obtained from the OMDB API
type Show struct {
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
	Ratings  []Rating
	Seasons  string `json:"totalSeasons"` // need to set this json value because it differs from the property name
}

// PrintShow takes the information stored in a Show struct and prints it out to the user
func (s Show) PrintShow() {
	fmt.Println("Title:", s.Title)
	fmt.Println("Year:", s.Year)
	fmt.Println("Rated:", s.Rated)
	fmt.Println("Released:", s.Released)
	fmt.Println("Runtime:", s.Runtime)
	fmt.Println("Genre:", s.Genre)
	fmt.Println("Director:", s.Director)
	fmt.Println("Writer:", s.Writer)
	fmt.Println("Cast:", s.Actors)
	fmt.Println("Plot:", s.Plot)
	fmt.Println("Language:", s.Language)
	fmt.Println("Awards:", s.Awards)
	fmt.Println("Number of Seasons:", s.Seasons)

	if len(s.Ratings) > 0 {
		fmt.Println("Ratings:")
		for i := range s.Ratings {
			rating := "* " + s.Ratings[i].Source + ": " + s.Ratings[i].Value
			fmt.Println(rating)
		}
	} else {
		fmt.Println("Ratings: N/A")
	}
}
