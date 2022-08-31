package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var tweetCounter = 0

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/tweets", tweetHandler)

	//log.Println("Listen at port :8080")
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

func tweetHandler(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Println("Failed to read body:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	tweet := Tweet{}

	if err := json.Unmarshal(body, &tweet); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", tweet.Message, tweet.Location)

	payload, err := json.Marshal(response{
		ID: tweetCounter + 1,
	})

	if err != nil {
		log.Println("Failed to marshal:", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	tweetCounter++
	writer.Write(payload)
}
