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
		fmt.Println("help")
		return
	}

	// get user's requested command (-m, -s, -h)
	cmd := args[1]

	switch cmd {
	case "-m":
		cmdArgs := buildString(args[2:])
		mov := performRequest(cmdArgs)
		mov.PrintMovie()
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

func performRequest(name string) models.Movie {
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
