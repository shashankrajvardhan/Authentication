package main

import (
	"fmt"
	"net/http"
)

const (
	Username = "admin"
	Password = "password123"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", serveLoginHTML)
	http.HandleFunc("/login", handleLogin)
	fmt.Println("Server listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
