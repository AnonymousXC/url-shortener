package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func write(tinyID *string, url *string) {

	var data map[string]string
	readConfig(&data)

	data[*tinyID] = *url

	var byteData, marshalError = json.Marshal(data)

	if marshalError != nil {
		log.Fatalln("Error occured while marshling data.", marshalError)
	}

	var writeError = os.WriteFile("sites.json", byteData, os.ModeAppend)
	if writeError != nil {
		log.Fatalln("Error occured writing to file.", writeError)
	} else {
		fmt.Println("Saved")
	}
}
