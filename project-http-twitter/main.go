package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Your solution goes here. Good luck!
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	s := server.New()

	router.Get("/tweets", s.ListTweets)
	router.Post("/tweets", s.AddTweet)

	go spamTweets()

	//log.Println("Listen at port :8080")
	err := http.ListenAndServe(":8080", router)

	log.Fatal(err)
}

func spamTweets() {
	for {
		// Prepare payload
		payload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		// Marshal payload
		jsonPayload, err := json.Marshal(payload)

		if err != nil {
			fmt.Println("Failed to marshal payload:", err)
			return
		}

		url := "http://localhost:8080/tweets"

		// Send HTTP POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))

		if err != nil {
			fmt.Println("failed to send http post:", err)
			return
		}

		// Close response body
		defer resp.Body.Close()

		// (Optionally read and print the response)
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("failed to read resp body:", err)
			return
		}

		var r = server.Response{}

		err = json.Unmarshal(body, &r)

		if err != nil {
			fmt.Println("failed to unmarshal body:", err)
		}

		fmt.Printf("response body %d", r.ID)
	}
}
