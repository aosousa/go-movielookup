package models

import (
	"fmt"
	"strconv"
)

// Season struct contains all the information of a TV show's season obtained from the OMDB API
type Season struct {
	Title    string
	Season   string
	Episodes BasicEpisodes
}

// PrintSeason takes the information stored in a Season struct and prints it out to the user
// Receives:
// * year (string) - Year of a show's season (if it was sent in the command line arguments)
func (s Season) PrintSeason(year string) {
	avgRating := 0.0
	var titleString string

	if year == "" {
		titleString = s.Title + " Season " + s.Season
	} else {
		titleString = s.Title + " (" + year + ") Season " + s.Season
	}

	fmt.Println(titleString)
	fmt.Println("Episodes:")
	for _, episode := range s.Episodes {
		rating, _ := strconv.ParseFloat(episode.Rating, 64)
		avgRating += rating

		fmt.Printf("* Episode %s: %s - %.1f IMDB rating\n", episode.Episode, episode.Title, rating)
	}
	avgRating = avgRating / float64(len(s.Episodes))
	fmt.Printf("Average Episode IMDB Rating: %.1f\n", avgRating)
}
