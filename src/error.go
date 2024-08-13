package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/error.html")
}
