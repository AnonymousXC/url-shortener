package main

import "net/http"

func handleSuccess(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/success.html")
}
