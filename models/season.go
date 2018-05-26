package models

import "fmt"

// Season struct contains all the information of a TV show's season obtained from the OMDB API
type Season struct {
	Title    string
	Season   string
	Episodes []BasicEpisode
}

// PrintSeason takes the information stored in a Season struct and prints it out to the user
func (s Season) PrintSeason() {
	fmt.Println(s.Title, "Season", s.Season)
	fmt.Println("Episodes:")
	for i := range s.Episodes {
		episode := "* Episode " + s.Episodes[i].Episode + ": " + s.Episodes[i].Title + " - " + s.Episodes[i].Rating + " IMDB rating"
		fmt.Println(episode)
	}
}
