package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

import "twitter/server"

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/tweets", s.ListTweets)
	router.Post("/tweets", s.AddTweet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
