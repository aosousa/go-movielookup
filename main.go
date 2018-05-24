package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	// get user's requested command
	// -m or -movie: Search for a movie
	// -s or -show: Search for a TV show
	// -h or -help: Print list of accepted commands
	cmd := args[1]

	switch cmd {
	case "-m", "-movie":
		cmdArgs := buildString(args[2:])
		mov := findMovie(cmdArgs)
		mov.PrintMovie()
	case "-s", "-show":
		cmdArgs := buildString(args[2:])
		show := findShow(cmdArgs)
		show.PrintShow()
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
	} else {
		defer res.Body.Close()
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		json.Unmarshal(content, &movie)
	}
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
	} else {
		defer res.Body.Close()
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		json.Unmarshal(content, &show)
	}
	return show
}

// Prints the list of accepted commands
func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("* -h | -help: Prints the list of available commands")
	fmt.Println("* -m | -movie `movie title`: Search for a movie (e.g. go-movie-lookup -m Avengers)")
	fmt.Println("* -s | -show `show title`: Search for a TV show (e.g. go-movie-lookup -s Game of Thrones)")
	fmt.Println("You can also search for a TV show season (e.g. go-movie-lookup -s Game of Thrones S3) or a TV show episode (e.g. go-movie-lookup -s Game of Thrones S3E5)")
}
