package main

import (
	"log"
	"net/http"
)

import "twitter/server"

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.Tweets)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
