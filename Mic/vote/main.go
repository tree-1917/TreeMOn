package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type VoteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Animal   string `json:"animal"` // fixed typo
}

var Votes = map[string]int{
	"cat": 0,
	"dog": 0,
}

func callAuth(username, password string) bool {
	authBody, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	resp, err := http.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(authBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	var req VoteRequest
	json.NewDecoder(r.Body).Decode(&req)

	if !callAuth(req.Username, req.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}

	if req.Animal != "cat" && req.Animal != "dog" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid vote"})
		return
	}

	Votes[req.Animal]++
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Votes)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Votes)
}

func main() {
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/results", resultsHandler) // fixed typo
	log.Println("Vote Service is Running On Port 8001...")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
