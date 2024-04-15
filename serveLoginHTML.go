package main

import "net/http"

func serveLoginHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}
