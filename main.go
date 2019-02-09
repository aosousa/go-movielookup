package main

import (
	"os"

	"github.com/aosousa/go-movielookup/utils"
)

func main() {
	// set up Config struct
	utils.InitConfig()

	// return help message if length of command line arguments is 1
	// we pick 1 instead of 0 because it'll never be 0
	// since os.Args[0] returns the name of the file being executed (in this case, go-movie-lookup.exe)
	args := os.Args
	if len(args) == 1 {
		utils.PrintHelp()
		return
	}

	cmd := args[1]

	switch cmd {
	case "-m", "--movie":
		utils.HandleMovie(args)
	case "-s", "--show":
		utils.HandleShowOptions(args)
	case "-h", "--help":
		utils.PrintHelp()
	case "-v", "--version":
		utils.PrintVersion()
	default:
		utils.PrintHelp()
	}
}
