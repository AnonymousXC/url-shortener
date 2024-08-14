package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "127.0.0.1:8080"

var SITE_MAP = make(map[string]string)

func main() {

	fmt.Println("Starting server on port : " + PORT)

	var mux = http.NewServeMux()
	mux.HandleFunc("/", Redirect)
	mux.HandleFunc("/add/new", handleAddNewSite)
	mux.HandleFunc("/add/success", handleSuccess)
	mux.HandleFunc("/add/error", handleError)

	log.Fatal(http.ListenAndServe(PORT, mux))
}
