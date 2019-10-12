package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	utils "github.com/aosousa/golang-utils"
)

// Config struct contains all the necessary configurations for the app
// to work correctly.
type Config struct {
	APIKey string `json:"apiKey"`
}

// CreateConfig adds information to a Config struct
func CreateConfig() Config {
	var config Config

	if _, err := os.Stat("./config.json"); err == nil {
		jsonFile, err := ioutil.ReadFile("./config.json")
		if err != nil {
			utils.HandleError(err)
		}

		err = json.Unmarshal(jsonFile, &config)
		if err != nil {
			utils.HandleError(err)
		}
	}

	config.APIKey = os.Getenv("OMDB_KEY")

	if config.APIKey == "" {
		err := errors.New("ERROR: OMDB API key missing")
		utils.HandleError(err)
	}

	return config
}
