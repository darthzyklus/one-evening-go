package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TweetsRepository interface {
	AddTweet(t Tweet) (int, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (repository *TweetsMemoryRepository) AddTweet(t Tweet) (int, error) {
	repository.tweets = append(repository.tweets, t)

	return len(repository.tweets), nil
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

	payload, err := json.Marshal(response{
		ID: id,
	})

	if err != nil {
		log.Println("Failed to marshal:", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	writer.Write(payload)
}

func main() {
	// Your solution goes here. Good luck!
	s := server{
		tweetsRepository: &TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.addTweet)

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
