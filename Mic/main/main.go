package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Structs
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VoteRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Animal   string `json:"animal"`
}

// -----------------
// Helpers
// -----------------

// Helper to call Auth service
func callAuth(loginReq LoginRequest) (*http.Response, error) {
	body, _ := json.Marshal(loginReq)
	return http.Post("http://localhost:8000/login", "application/json", bytes.NewBuffer(body))
}

// Helper to call Vote service
func callVote(voteReq VoteRequest) (*http.Response, error) {
	body, _ := json.Marshal(voteReq)
	return http.Post("http://localhost:8001/vote", "application/json", bytes.NewBuffer(body))
}

// Helper to proxy response from another service
func proxyResponse(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

// -----------------
// Handlers
// -----------------

// WrapHandler adds logging of duration and method
func WrapHandler(name string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s processed in %v", name, r.Method, r.URL.Path, duration)
	}
}

// /main_login
func mainLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := callAuth(req)
	if err != nil {
		http.Error(w, "Auth service error", http.StatusInternalServerError)
		return
	}
	proxyResponse(w, resp)
}

// /main_vote
func mainVoteHandler(w http.ResponseWriter, r *http.Request) {
	var req VoteRequest
	json.NewDecoder(r.Body).Decode(&req)

	resp, err := callVote(req)
	if err != nil {
		http.Error(w, "Vote service error", http.StatusInternalServerError)
		return
	}
	proxyResponse(w, resp)
}

// /main_results
func mainResultsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8001/results")
	if err != nil {
		http.Error(w, "Vote service error", http.StatusInternalServerError)
		return
	}
	proxyResponse(w, resp)
}

// -----------------
// Main
// -----------------

func main() {
	http.HandleFunc("/main_login", WrapHandler("MAIN_LOGIN", mainLoginHandler))
	http.HandleFunc("/main_vote", WrapHandler("MAIN_VOTE", mainVoteHandler))
	http.HandleFunc("/main_results", WrapHandler("MAIN_RESULTS", mainResultsHandler))

	fmt.Println("Main gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
