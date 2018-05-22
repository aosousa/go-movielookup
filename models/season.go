package models

// Season struct contains all the information of a TV show's season obtained from the OMDB API
type Season struct {
	Title   string
	Season  int
	Episode []BasicEpisode
}
