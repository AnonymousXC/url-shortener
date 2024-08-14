package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func readConfig(data *map[string]string) {

	var file, err = os.ReadFile("sites.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.WriteFile("sites.json", []byte("{}"), 0644)
			return
		}
		log.Fatalln("Error occured reading the file.", err)
	}

	var unmarshError = json.Unmarshal(file, &data)

	if unmarshError != nil {
		log.Fatalln("Error occured parsing the file.", unmarshError)
	}
}
