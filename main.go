package main

import (
	"fmt"
	"os"
)

func main() {
	apiKey := "a5fafd94"
	baseURL := "http://www.omdbapi.com/?apikey=" + apiKey + "&"

	// return help message if length of command line arguments is 1
	// we pick 1 instead of 0 because it'll never be 0
	// since os.Args[0] returns the name of the file being executed (in this case, movie-lookup.exe)
	args := os.Args
	if len(args) == 1 {
		fmt.Println("help")
		return
	}

	// os.Args[0] returns the name of the file being executed (in this case, movie-lookup.exe)
	cmd := args[1]

	switch cmd {
	case "-m":
		GetMovie(args[2])
	}
	//fmt.Println(cmd)

	fmt.Println(baseURL)
}

func test() {
	fmt.Println("searched for movie")
}
