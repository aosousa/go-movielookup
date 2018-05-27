package main

import (
	"os"
)

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
	case "-m", "--movie":
		handleMovie(args)
	case "-s", "--show":
		handleShowOptions(args)
	case "-h", "--help":
		printHelp()
	case "-v", "--version":
		printVersion()
	default:
		printHelp()
	}
}
