package models

// Error struct contains the information in case of an error from the OMDB API
type Error struct {
	Response bool
	Error    string
}
