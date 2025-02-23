package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/tweets", s.ListTweets)
	router.Post("/tweets", s.AddTweet)

	go spamTweets()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func spamTweets() {
	for {
		time.Sleep(time.Millisecond * 100)

		// Prepare payload
		payload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		// Marshal payload
		payloadJSON, err := json.Marshal(payload)

		if err != nil {
			fmt.Printf("Failed to marshall payload: %v", err)
			return
		}

		// Send HTTP POST Request
		url := "http://localhost:8080/tweets"
		contentType := "application/json"

		resp, err := http.Post(url, contentType, bytes.NewBuffer(payloadJSON))

		if err != nil {
			fmt.Printf("Error at the HTTP POST request: %v", err)
			return
		}

		// Defer close of the response  body
		defer resp.Body.Close()

		// read and print the response

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Failed to read the response body: %v", err)
			return
		}

		respBody := server.TweetResponse{}

		if err := json.Unmarshal(body, &respBody); err != nil {
			fmt.Printf("Failed to unmarshal response body: %v", err)
			return
		} else {
			fmt.Printf("Added tweet %d\n", respBody.ID)
		}

	}
}
