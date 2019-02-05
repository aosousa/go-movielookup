package models

import "fmt"

// Show struct contains all the information of a TV show obtained from the OMDB API
type Show struct {
	OMDBObject
	Seasons string `json:"totalSeasons"` // need to set this json value because it differs from the property name
}

// PrintShow takes the information stored in a Show struct and prints it out to the user
func (s Show) PrintShow() {
	s.OMDBObject.PrintOMDBInfo()
	fmt.Println("Number of Seasons:", s.Seasons)
	s.OMDBObject.Ratings.PrintRatings()
}
