package models

import "fmt"

// Error struct contains the information in case of an error from the OMDB API
type Error struct {
	Response string
	Error    string
}

// PrintError takes the error information sent by the API and prints it out to the user
func (e Error) PrintError() {
	fmt.Println("Error:", e.Error)
}
