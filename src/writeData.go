package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func write(tinyID *string, url *string) {

	var file, readError = os.ReadFile("sites.json")
	if readError != nil {
		log.Fatalln("Error occured.", readError)
	}

	var data map[string]string
	var unmarshError = json.Unmarshal(file, &data)

	if unmarshError != nil {
		log.Fatalln("Error occured.", unmarshError)
	}

	data[*tinyID] = *url

	var byteData, marshalError = json.Marshal(data)

	if marshalError != nil {
		log.Fatalln("Error occured", marshalError)
	}

	var writeError = os.WriteFile("sites.json", byteData, os.ModeAppend)
	if writeError != nil {
		log.Fatalln("Error occured", writeError)
	} else {
		fmt.Println("Saved")
	}
}
