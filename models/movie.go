package models

import "fmt"

// Movie struct contains all the information of a movie obtained
// from the OMDB API
type Movie struct {
	OMDBObject
	Production string
}

// PrintMovie takes the information stored in a Movie struct and prints it out to the user
func (m Movie) PrintMovie() {
	m.OMDBObject.PrintOMDBInfo()
	fmt.Println("Production:", m.Production)
	m.OMDBObject.Ratings.PrintRatings()
}
