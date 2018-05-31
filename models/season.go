package models

import "fmt"

// Season struct contains all the information of a TV show's season obtained from the OMDB API
type Season struct {
	Title    string
	Season   string
	Episodes []BasicEpisode
}

// PrintSeason takes the information stored in a Season struct and prints it out to the user
// Receives:
// * year (string) - Year of a show's season (if it was sent in the command line arguments)
func (s Season) PrintSeason(year string) {
	var titleString string
	if year == "" {
		titleString = s.Title + " Season " + s.Season
	} else {
		titleString = s.Title + " (" + year + ") Season " + s.Season
	}
	fmt.Println(titleString)
	fmt.Println("Episodes:")
	for i := range s.Episodes {
		episode := "* Episode " + s.Episodes[i].Episode + ": " + s.Episodes[i].Title + " - " + s.Episodes[i].Rating + " IMDB rating"
		fmt.Println(episode)
	}
}
