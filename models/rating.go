package models

// Rating struct contains all the information of a movie or TV show's
// rating (IMDb, Rotten Tomatoes, Metacritic) obtained from the OMDB API
type Rating struct {
	Source string
	Value  string
}
