package main

import (
	"net/http"
)

func handleRedirectRequest(w http.ResponseWriter, r *http.Request) {

	var URL_ID = r.URL.Path[1:]
	var SITE_URL, exists = SITE_MAP[URL_ID]

	if URL_ID == "" {
		handleAddServe(w, r)
	}

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, SITE_URL, http.StatusSeeOther)
}
