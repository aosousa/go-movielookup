package models

// Show struct contains all the information of a TV show obtained from the OMDB API
type Show struct {
	Title    string
	Year     int
	Rated    string
	Released string
	Runtime  string
	Genre    string
	Director string
	Writer   string
	Cast     string
	Plot     string
	Language string
	Awards   string
	Ratings  []Rating
	Seasons  int
}
