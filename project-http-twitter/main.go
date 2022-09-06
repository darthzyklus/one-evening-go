package main

import (
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

	//log.Println("Listen at port :8080")
	err := http.ListenAndServe(":8080", router)

	log.Fatal(err)
}
