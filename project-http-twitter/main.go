package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TweetsRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (repository *TweetsMemoryRepository) AddTweet(t Tweet) (int, error) {
	repository.tweets = append(repository.tweets, t)

	return len(repository.tweets), nil
}

func (repository *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	tweets := repository.tweets

	return tweets, nil
}

type server struct {
	tweetsRepository TweetsRepository
}

func (s server) addTweet(writer http.ResponseWriter, req *http.Request) {
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

	id, err := s.tweetsRepository.AddTweet(tweet)

	if err != nil {
		log.Println("Failed to save the tweet:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var resp = response{ID: id}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		log.Println("Failed to marshal:", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(respJSON)
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s server) listTweets(writer http.ResponseWriter, req *http.Request) {
	tweets, err := s.tweetsRepository.Tweets()

	if err != nil {
		log.Println("Failed to get tweets:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var resp = tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		log.Println("Failed to marshal:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(respJSON)
}

func (s server) tweets(writer http.ResponseWriter, req *http.Request) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", req.Method, req.URL, duration)
	}()

	if req.Method == http.MethodPost {
		s.addTweet(writer, req)
	} else if req.Method == http.MethodGet {
		s.listTweets(writer, req)
	}
}

func main() {
	// Your solution goes here. Good luck!
	s := server{
		tweetsRepository: &TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)

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
