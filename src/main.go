package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8000"

var SITE_MAP = make(map[string]string)

func main() {

	fmt.Println("Starting server on port : " + PORT)
	readConfig(&SITE_MAP)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", Redirect)
	mux.HandleFunc("/add/new", handleAddNewSite)
	mux.HandleFunc("/add/success", handleSuccess)
	mux.HandleFunc("/add/error", handleError)
	mux.HandleFunc("/api", api)

	log.Fatal(http.ListenAndServe(PORT, mux))
}
