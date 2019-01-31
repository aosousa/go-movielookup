package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct contains all the necessary configurations for the app
// to work correctly.
type Config struct {
	APIKey string `json:"apiKey"`
}

// CreateConfig adds information to a Config struct
func CreateConfig() Config {
	var config Config
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return config
}
