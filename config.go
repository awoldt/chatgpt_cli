package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Key   string `json:"key"`
	Model string `json:"model"`
}

var UserConfig Config

func init() {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		// see if the error is about the config file not being present
		if errors.Is(err, os.ErrNotExist) {
			josnData, err := json.Marshal(Config{Key: "", Model: "o4-mini"})
			if err != nil {
				log.Fatal("error: could not generate config.json in root")
			}

			os.WriteFile("config.json", josnData, 0666)
			fmt.Println("Generated a config file in root. Add openai api key to start using.")
			os.Exit(0)
		} else {
			log.Fatal("error: could not generate config.json in root")
		}
	}

	err = json.Unmarshal(data, &UserConfig)
	if err != nil {
		log.Fatal("error: could not establish config details")
	}

	if UserConfig.Key == "" {
		log.Fatal("error: must provide api key in the config.json")
	}
}
