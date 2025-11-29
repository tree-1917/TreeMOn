package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// user struct
type User struct {
	Username string
	Password string
}

// in-memory database
var users = []User{
	{Username: "admin", Password: "1234"},
	{Username: "moussa", Password: "1917"},
}

// Core Auth logic
func authCore(username, password string) bool {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false

}

// Schema fro login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if authCore(req.Username, req.Password) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
}

// Main func
func main() {
	http.HandleFunc("/login", loginHandler)
	log.Println("Auth Service running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
