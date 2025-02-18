package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var tweetsCounter = 0

type TweetResponse struct {
	ID int
}

type Tweet struct {
	ID       int    `json:"ID"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

func addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	tw := Tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		fmt.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tweetsCounter++
	tw.ID = tweetsCounter

	resp := TweetResponse{ID: tw.ID}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

func main() {
	http.HandleFunc("/tweets", addTweet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
