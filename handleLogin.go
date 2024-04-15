package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var login Login

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var err error
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if login.Username == Username && login.Password == Password {
		fmt.Println("Authentication successful")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	} else {
		fmt.Println("Authentication failed")

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Login failed"))
	}

}
