package models

// Episode struct contains all the information of a TV show episode obtained from the OMDB API
type Episode struct {
	Title    string
	Rated    string
	Released string
	Season   int
	Number   int
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Plot     string
	Ratings  []Rating
}
