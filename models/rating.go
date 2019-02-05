package models

import "fmt"

// Rating struct contains all the information of a movie or TV show's
// rating (IMDb, Rotten Tomatoes, Metacritic) obtained from the OMDB API
type Rating struct {
	Source string
	Value  string
}

// Ratings represents a slice of Rating structs
type Ratings []Rating

// PrintRatings takes the information store in a Ratings struct and prints it out to the user
func (r Ratings) PrintRatings() {
	if len(r) > 0 {
		fmt.Println("Ratings:")
		for _, rating := range r {
			fmt.Printf("* %s: %s\n", rating.Source, rating.Value)
		}
	} else {
		fmt.Println("Ratings: N/A")
	}
}
