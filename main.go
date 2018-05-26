package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"./models"
)

var apiKey = "a5fafd94"
var baseURL = "http://www.omdbapi.com/?apikey=" + apiKey + "&"

func main() {
	// return help message if length of command line arguments is 1
	// we pick 1 instead of 0 because it'll never be 0
	// since os.Args[0] returns the name of the file being executed (in this case, go-movie-lookup.exe)
	args := os.Args
	if len(args) == 1 {
		printHelp()
		return
	}

	cmd := args[1]

	switch cmd {
	case "-m", "-movie":
		cmdArgs := buildString(args[2:])
		mov := findMovie(cmdArgs)
		mov.PrintMovie()
	case "-s", "-show":
		lastArg := args[len(args)-1]
		nextToLastArg := args[len(args)-2]

		// get last cmd line argument to check if episode argument was sent
		episodeRegex, _ := regexp.Compile("E([0-9]+)")
		episodeRegexArgs := episodeRegex.FindStringSubmatch(lastArg)

		// episode argument was found, check for a season argument
		// else, check for a season argument again but to print only season information
		if len(episodeRegexArgs) > 0 {
			seasonRegex, _ := regexp.Compile("S([0-9]+)")
			seasonRegexArgs := seasonRegex.FindStringSubmatch(nextToLastArg)

			// season was found, print episode information
			// else print error message stating that season number is required for this functionality
			if len(seasonRegexArgs) > 0 {
				showTitle := buildString(args[2 : len(args)-2])
				seasonNumber := seasonRegexArgs[1]
				episodeNumber := episodeRegexArgs[1]

				episode := findEpisode(showTitle, seasonNumber, episodeNumber)
				episode.PrintEpisode()
			} else {
				fmt.Println("err")
			}
		} else {
			seasonRegex, _ := regexp.Compile("S([0-9]+)")
			seasonRegexArgs := seasonRegex.FindStringSubmatch(lastArg)

			// season was found, print season information
			// else print show information
			if len(seasonRegexArgs) > 0 {
				showTitle := buildString(args[2 : len(args)-1])
				seasonNumber := seasonRegexArgs[1]

				season := findSeason(showTitle, seasonNumber)
				season.PrintSeason()
			} else {
				cmdArgs := buildString(args[2:])
				show := findShow(cmdArgs)
				show.PrintShow()
			}
		}
	case "-h", "-help":
		printHelp()
	}
}

// Builds a string that can be used in the API request in case the movie or show has more than word
// e.g. Avengers Infinity War, Game of Thrones
func buildString(args []string) string {
	var name string
	for _, v := range args {
		name += v + "+"
	}

	nameLen := len(name)

	// need to remove the last + at the end, otherwise it might return the wrong movie or show in some situations
	// e.g. Avengers+ returns Avengers: Age of Ultron, Avengers returns the first movie
	return name[:nameLen-1]
}

// Performs an HTTP request to find a movie with the title provided by the user
func findMovie(name string) models.Movie {
	movie := models.Movie{}

	queryURL := baseURL + "t=" + name
	res, err := http.Get(queryURL)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	json.Unmarshal(content, &movie)

	return movie
}

// Performs an HTTP request to find a TV show with the title provided by the user
func findShow(name string) models.Show {
	show := models.Show{}

	queryURL := baseURL + "t=" + name + "&type=series"
	res, err := http.Get(queryURL)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	json.Unmarshal(content, &show)

	return show
}

// Performs an HTTP request to find a TV show episode with the title, season, and episode number provided by the user
func findEpisode(show, season, epNumber string) models.Episode {
	episode := models.Episode{}

	queryURL := baseURL + "t=" + show + "&type=series&season=" + season + "&episode=" + epNumber

	res, err := http.Get(queryURL)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	json.Unmarshal(content, &episode)

	return episode
}

// Performs an HTTP request to find a TV show season with the title and season number provided by the user
func findSeason(title, seasonNumber string) models.Season {
	season := models.Season{}

	queryURL := baseURL + "t=" + title + "&type=series&season=" + seasonNumber

	res, err := http.Get(queryURL)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	json.Unmarshal(content, &season)

	return season
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Println("Movie and TV Show Lookup")
	fmt.Println("Available commands:")
	fmt.Println("* -h | -help: Prints the list of available commands")
	fmt.Println("* -m | -movie `movie title`: Search for a movie (e.g. go-movie-lookup -m Avengers)")
	fmt.Println("* -s | -show `show title`: Search for a TV show (e.g. go-movie-lookup -s Game of Thrones)")
	fmt.Println("You can also search for a TV show season (e.g. go-movie-lookup -s Game of Thrones S3) or a TV show episode (e.g. go-movie-lookup -s Game of Thrones S3 E5)")
}
