package main

import (
	"encoding/json"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// if r.Method == "GET" {
	// 	fmt.Fprint(w, "POST request required.")
	// 	return
	// }

	var params = r.URL.Query()
	if !params.Has("id") || !params.Has("url") {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  400,
			"message": "Parameters not found.",
		})
		return
	}

	var tinyID = params["id"][0]
	var _, doesExists = SITE_MAP[tinyID]

	if doesExists {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  409,
			"message": "Id already exists.",
		})
		return
	} else {
		SITE_MAP[tinyID] = params["url"][0]
		write(&tinyID, &params["url"][0])
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  200,
			"message": "Successfully added to DB.",
			"url":     "https://" + r.Host + "/" + tinyID,
		})
	}
}
