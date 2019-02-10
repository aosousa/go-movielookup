package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/aosousa/go-movielookup/models"
)

const version = "1.3.0"

var (
	baseURL string
	config  models.Config
)

// InitConfig reads the configuration file and parses that information into a Config struct
func InitConfig() {
	config = models.CreateConfig()
	baseURL = "http://www.omdbapi.com/?apikey=" + config.APIKey + "&"
}

/*HandleMovie handles a request to lookup a movie with the title provided by the user
 * Receives:
 * args ([]string) - Arguments passed in the terminal by the user
 */
func HandleMovie(args []string) {
	var queryURL string
	apiError := models.Error{}

	lastArg := args[len(args)-1]
	yearRegex, _ := regexp.Compile(`\([0-9]+\)`)
	yearRegexArg := yearRegex.FindStringSubmatch(lastArg)

	// if year was not sent in the command line arguments, perform normal movie query
	// if it was, add year to the query URL
	if len(yearRegexArg) == 0 {
		movieTitle := buildTitleString(args[2:])
		queryURL = baseURL + "t=" + movieTitle + "&type=movie"
	} else {
		// build title string but remove year argument
		movieTitle := buildTitleString(args[2 : len(args)-1])
		queryURL = baseURL + "t=" + movieTitle + "&type=movie&y=" + yearRegexArg[0][1:5]
	}

	res, err := http.Get(queryURL)
	checkResponse(res.StatusCode)

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

	// if no error occurred, print the movie information
	// else, print error message sent in the API
	json.Unmarshal(content, &apiError)
	if apiError.Response == "True" {
		movie := models.Movie{}
		json.Unmarshal(content, &movie)
		movie.PrintMovie()
	} else {
		apiError.PrintError()
	}
}

/*HandleShowOptions handles a request to lookup a TV show, TV show season, or TV show episode
 * Receives:
 * args ([]string) - Arguments passed in the terminal by the user
 */
func HandleShowOptions(args []string) {
	lastArg := args[len(args)-1]
	nextToLastArg := args[len(args)-2]

	// get last cmd line argument to check if episode argument was sent
	episodeRegex, _ := regexp.Compile("(?i)E([0-9]+)")
	episodeRegexArgs := episodeRegex.FindStringSubmatch(lastArg)

	// if episode argument was found, check for a season argument
	// else, check for a season argument again but to print only season information
	if len(episodeRegexArgs) > 0 {
		seasonRegex, _ := regexp.Compile("(?i)S([0-9]+)")
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
		seasonRegex, _ := regexp.Compile("(?i)S([0-9]+)")
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
	var queryURL string
	apiError := models.Error{}

	lastArg := args[len(args)-1]
	yearRegex, _ := regexp.Compile(`\([0-9]+\)`)
	yearRegexArg := yearRegex.FindStringSubmatch(lastArg)

	// if year was not sent in the command line arguments, perform normal show query
	// if it was, add year to the query URL
	if len(yearRegexArg) == 0 {
		showTitle := buildTitleString(args[2:])
		queryURL = baseURL + "t=" + showTitle + "&type=series"
	} else {
		// build title string but remove year argument
		movieTitle := buildTitleString(args[2 : len(args)-1])
		queryURL = baseURL + "t=" + movieTitle + "&type=series&y=" + yearRegexArg[0][1:5]
	}

	res, err := http.Get(queryURL)
	checkResponse(res.StatusCode)

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

	// if no error occurred, print show information
	// else, print error message sent in the APi
	json.Unmarshal(content, &apiError)
	if apiError.Response == "True" {
		show := models.Show{}
		json.Unmarshal(content, &show)
		show.PrintShow()
	} else {
		apiError.PrintError()
	}
}

// Handles a request to lookup a TV show season
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
// * seasonNum (string) - Season number
func findSeason(args []string, seasonNum string) {
	var queryURL string
	var year string
	apiError := models.Error{}

	// next to last argument here because the last argument is the season number
	lastArg := args[len(args)-2]
	yearRegex, _ := regexp.Compile(`\([0-9]+\)`)
	yearRegexArg := yearRegex.FindStringSubmatch(lastArg)

	// if year was not sent in the command line arguments, perform normal season query
	// if it was, add year to the query URL
	if len(yearRegexArg) == 0 {
		year = ""

		// build title string but remove season number (last argument)
		showTitle := buildTitleString(args[2 : len(args)-1])
		queryURL = baseURL + "t=" + showTitle + "&type=series&season=" + seasonNum
	} else {
		// build title string but remove season number AND year (last 2 arguments)
		year = yearRegexArg[0][1:5]
		showTitle := buildTitleString(args[2 : len(args)-2])
		queryURL = baseURL + "t=" + showTitle + "&type=series&season=" + seasonNum + "&y=" + year
	}

	res, err := http.Get(queryURL)
	checkResponse(res.StatusCode)

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

	// if no error occurred, print season information
	// else, print error message sent in the API
	json.Unmarshal(content, &apiError)
	if apiError.Response == "True" {
		season := models.Season{}
		json.Unmarshal(content, &season)
		season.PrintSeason(year)
	} else {
		apiError.PrintError()
	}
}

// Handles a request to lookup a TV show episode
// Receives:
// * args ([]string) - Arguments passed in the terminal by the user
// * season (string) - Season number
// * episodeNum (string) - Episode number
func findEpisode(args []string, season string, episodeNum string) {
	var queryURL string
	apiError := models.Error{}

	// next to next to last argument here since the last 2 arguments are season and episode number
	lastArg := args[len(args)-3]
	yearRegex, _ := regexp.Compile(`\([0-9]+\)`)
	yearRegexArg := yearRegex.FindStringSubmatch(lastArg)

	// if year was not sent in the command line arguments, perform normal episode query
	// if it was, add year to the query URL

	if len(yearRegexArg) == 0 {
		// build title string but remove season and episode numbers (last 2 arguments)
		showTitle := buildTitleString(args[2 : len(args)-2])
		queryURL = baseURL + "t=" + showTitle + "&type=series&season=" + season + "&episode=" + episodeNum
	} else {
		// build title string but remove season, episode, AND year (last 3 arguments)
		showTitle := buildTitleString(args[2 : len(args)-3])
		queryURL = baseURL + "t=" + showTitle + "&type=series&season=" + season + "&episode=" + episodeNum + "&y=" + yearRegexArg[0][1:5]
	}

	res, err := http.Get(queryURL)
	checkResponse(res.StatusCode)

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

	// if no error occurred, print episode information
	// else, print error message sent in the API
	json.Unmarshal(content, &apiError)
	if apiError.Response == "True" {
		episode := models.Episode{}
		json.Unmarshal(content, &episode)
		episode.PrintEpisode()
	} else {
		apiError.PrintError()
	}
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

// PrintHelp prints the list of accepted commands
func PrintHelp() {
	fmt.Println("Movie and TV Show Lookup (version " + version + ")")
	fmt.Println("Available commands:")
	fmt.Println("* -h | --help    Prints the list of available commands")
	fmt.Println("* -v | --version Prints the version of the program")
	fmt.Println("\n* -m | --movie `movie title` [(YEAR)] Search for a movie (e.g. go-movie-lookup -m Avengers)")
	fmt.Println("In case you want the movie from a specific year, you can add the year in front of the movie title (e.g. go-movie-lookup -m Ghostbusters (1984)")
	fmt.Println("\n* -s | --show `show title` [S1 | S1 E1] Search for a TV show (e.g. go-movie-lookup -s Game of Thrones)")
	fmt.Println("You can also search for a TV show season (e.g. go-movie-lookup -s Game of Thrones S3) or a TV show episode (e.g. go-movie-lookup -s Game of Thrones S3 E9)")
	fmt.Println("In case you want the TV show from a specific year, you can add the year in front of the show title (e.g. go-movie-lookup -s House of Cards (1990)")
}

// PrintVersion prints the version of the program
func PrintVersion() {
	fmt.Println("Version " + version)
}

// Prints an error stating the correct format to look up for a TV show episode
func printShowFormatError() {
	fmt.Println("Error: The correct format to look up for a TV show episode is:")
	fmt.Println("go-movielookup -s (Show title) S(number) E(number) (e.g. go-movielookup -s Game of Thrones S3 E9)")
}

// Check if there was an error with the API
// Receives:
// * code (int) - Response status code
func checkResponse(code int) {
	if code != 200 {
		fmt.Println("Error: An error occurred while performing your request. Please try again later.")
		os.Exit(1)
	}
}
