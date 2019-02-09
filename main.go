package main

import (
	"os"

	"github.com/aosousa/go-movielookup/utils"
)

func main() {
	// set up Config struct
	utils.initConfig()

	// return help message if length of command line arguments is 1
	// we pick 1 instead of 0 because it'll never be 0
	// since os.Args[0] returns the name of the file being executed (in this case, go-movie-lookup.exe)
	args := os.Args
	if len(args) == 1 {
		utils.printHelp()
		return
	}

	cmd := args[1]

	switch cmd {
	case "-m", "--movie":
		utils.handleMovie(args)
	case "-s", "--show":
		utils.handleShowOptions(args)
	case "-h", "--help":
		utils.printHelp()
	case "-v", "--version":
		utils.printVersion()
	default:
		utils.printHelp()
	}
}
