package models

// BasicEpisode struct contains all the information of a TV show's episode obtained
// from the OMDB API to use in the season struct
type BasicEpisode struct {
	Title    string
	Number   int
	Released string
	Rating   string
}