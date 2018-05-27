package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"go-movielookup/models"
)

// var apiKey = "a5fafd94"
// var baseURL = "http://www.omdbapi.com/?apikey=" + apiKey + "&"
var version = "0.1.0"

// Handles a request to lookup a movie with the title provided by the user
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
func handleMovie(args []string) {
	cmdArgs := buildTitleString(args[2:])
	movie := findMovie(cmdArgs)
	movie.PrintMovie()
}

// Performs an HTTP request to find a movie with the title provided by the user
// Receives:
// * name (string) - Name of the movie
func findMovie(name string) models.Movie {
	movie := models.Movie{}

	queryURL := baseURL + "t=" + name
	res, err := http.Get(queryURL)
	if err != nil {
		fmt.Println("%s", err)
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

// Handles a request to lookup a TV show, TV show season, or TV show episode
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
func handleShowOptions(args []string) {
	lastArg := args[len(args)-1]
	nextToLastArg := args[len(args)-2]

	// get last cmd line argument to check if episode argument was sent
	episodeRegex, _ := regexp.Compile("E([0-9]+)")
	episodeRegexArgs := episodeRegex.FindStringSubmatch(lastArg)

	// if episode argument was found, check for a season argument
	// else, check for a season argument again but to print only season information
	if len(episodeRegexArgs) > 0 {
		seasonRegex, _ := regexp.Compile("S([0-9]+)")
		seasonRegexArgs := seasonRegex.FindStringSubmatch(nextToLastArg)

		// if season was found, print season information
		// else print error message stating that season number is required for this functionality
		if len(seasonRegexArgs) > 0 {
			seasonNumber := seasonRegexArgs[1]
			episodeNumber := episodeRegexArgs[1]

			findEpisode(args, seasonNumber, episodeNumber)
		} else {
			printShowFormatError()
		}
	} else {
		seasonRegex, _ := regexp.Compile("S([0-9]+)")
		seasonRegexArgs := seasonRegex.FindStringSubmatch(lastArg)

		// if season was found, print season information
		// else print show information
		if len(seasonRegexArgs) > 0 {
			seasonNumber := seasonRegexArgs[1]

			findSeason(args, seasonNumber)
		} else {
			findShow(args)
		}
	}
}

// Handles a request to lookup a TV show
func findShow(args []string) {
	show := models.Show{}
	showTitle := buildTitleString(args[2:])

	queryURL := baseURL + "t=" + showTitle + "&type=series"
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

	show.PrintShow()
}

// Handles a request to lookup a TV show season
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
// * seasonNum (string) - Season number
func findSeason(args []string, seasonNum string) {
	season := models.Season{}
	showTitle := buildTitleString(args[2 : len(args)-2])

	queryURL := baseURL + "t=" + showTitle + "&type=series&season=" + seasonNum

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

	season.PrintSeason()
}

// Handles a request to lookup a TV show episode
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
// * season (string) - Season number
// * episodeNum (string) - Episode number
func findEpisode(args []string, season string, episodeNum string) {
	episode := models.Episode{}
	showTitle := buildTitleString(args[2 : len(args)-2])

	queryURL := baseURL + "t=" + showTitle + "&type=series&season=" + season + "&episode=" + episodeNum

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

	episode.PrintEpisode()
}

// Builds a string that can be used in the API request in case the movie or show title has more than one word
// e.g. Avengers Infinity War, Game of Thrones
func buildTitleString(args []string) string {
	var name string
	for _, v := range args {
		name += v + "+"
	}

	nameLen := len(name)

	// need to remove the last "+" at the end, otherwise it might return the wrong movie or show in some situations
	// e.g. "Avengers+"" returns Avengers: Age of Ultron, "Avengers" returns the first movie
	return name[:nameLen-1]
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Println("Movie and TV Show Lookup (version " + version + ")")
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help: Prints the list of available commands")
	fmt.Println("* -v | --version: Prints the version of the program")
	fmt.Println("* -m | --movie `movie title`: Search for a movie (e.g. go-movie-lookup -m Avengers)")
	fmt.Println("* -s | --show `show title`: Search for a TV show (e.g. go-movie-lookup -s Game of Thrones)")
	fmt.Println("You can also search for a TV show season (e.g. go-movie-lookup -s Game of Thrones S3) or a TV show episode (e.g. go-movie-lookup -s Game of Thrones S3 E9)")
}

// Prints the version of the program
func printVersion() {
	fmt.Println("Version " + version)
}

// Prints an error stating the correct format to look up for a TV show episode
func printShowFormatError() {
	fmt.Println("Error: The correct format to look up for a TV show episode is:")
	fmt.Println("go-movielookup -s (Show title) S(number) E(number) (e.g. go-movielookup -s Game of Thrones S3 E9")
}
