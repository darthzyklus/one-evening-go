package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	s := server{
		tweetsRepository: &TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.addTweet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetResponse struct {
	ID int
}

type TweetsRepository interface {
	SaveTweet(tw Tweet) (int, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (r *TweetsMemoryRepository) SaveTweet(tw Tweet) (int, error) {
	r.tweets = append(r.tweets, tw)

	return len(r.tweets), nil
}

type server struct {
	tweetsRepository TweetsRepository
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
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

	tweetID, err := s.tweetsRepository.SaveTweet(tw)

	if err != nil {
		fmt.Println("Failed to save tweet:", nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := TweetResponse{ID: tweetID}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}
