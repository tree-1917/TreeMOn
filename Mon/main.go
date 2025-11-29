package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"mon/auth"
	"mon/vote"
)

type VoteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Animal   string `json:"animal"`
}

// -----------------
// Handlers
// -----------------

func voteHandler(w http.ResponseWriter, r *http.Request) {
	var req VoteRequest
	json.NewDecoder(r.Body).Decode(&req)

	// 1. Check Auth
	if !auth.AuthCore(req.Username, req.Password) {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	// 2. Vote
	ok := vote.VoteCore(req.Animal)
	if !ok {
		http.Error(w, "Invaild vote ( cat or dog only)", http.StatusBadRequest)
		return
	}

	// 3. Return results
	json.NewEncoder(w).Encode(vote.Votes)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(vote.Votes)
}
func wrapHandler(name string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s processed in %v", name, r.Method, r.URL.Path, duration)
	}
}

// -----------------
// Main
// -----------------
func main() {
	fmt.Println("Monlithic app running on :8080 ...")

	http.HandleFunc("/vote", wrapHandler("Vote", voteHandler))
	http.HandleFunc("/results", wrapHandler("Result", resultsHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
