package main

import "net/http"

func handleAddServe(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/add.html")
}
