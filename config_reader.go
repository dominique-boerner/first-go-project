package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	i18nPath    string
	filePattern string
}

func ReadConfig() Config {
	fileContent, err := os.ReadFile("./config.json")

	if err != nil {
		log.Fatal("Error when opening config file: ", err)
		return Config{}
	}

	var payload Config
	decodeErr := json.NewDecoder(bytes.NewBuffer(fileContent)).Decode(&payload)
	if decodeErr != nil {
		log.Fatal("Error while decoding config file: ", err)
	}

	return payload

}
