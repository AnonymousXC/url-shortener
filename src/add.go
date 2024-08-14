package main

import (
	"net/http"
)

func handleAddNewSite(w http.ResponseWriter, r *http.Request) {
	var params = r.URL.Query()
	var tinyID = params["id"][0]
	var _, doesExists = SITE_MAP[tinyID]

	if doesExists {
		http.Redirect(w, r, "/add/error", http.StatusSeeOther)
		return
	} else {
		SITE_MAP[tinyID] = params["url"][0]
		http.Redirect(w, r, "/add/success?id="+tinyID, http.StatusSeeOther)
		write(&tinyID, &params["url"][0])
	}
}
