package main

import (
	"fmt"
	"net/http"
)

const successHTML = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet"
		        integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
		    <title>Shortly</title>
		</head>
		<body data-bs-theme="dark">
		    <div class="alert alert-success position-absolute top-50 start-50 w-75 translate-middle" role="alert" id="message">
		        Your link was added successfully.
		        <br>
		         <br>
		        Your new url is : <a id="url"></a> 
		        <br>
		        <button class="btn btn-dark mt-4" onclick="handleCopy()">Click to copy</button>
		      </div>
		    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"
		    integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz"
		    crossorigin="anonymous"></script>
		</body>
		<script>
		    const url = new URL(window.location.toString())
		    const id = url.searchParams.get("id")
		    if(id !== "") {
		        document.getElementById("url").innerHTML = url.origin + "/" + id
		        document.getElementById("url").href = url.origin + "/" + id
		    }

		    function handleCopy() {
		        if(id === "") return;
		        window.navigator.clipboard.writeText(url.origin + "/" + id)
		        alert("Successfully copied.")
		    }
		</script>
		</html>
`

func handleSuccess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, successHTML)
}
