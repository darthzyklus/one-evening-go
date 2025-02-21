package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	s := server{
		tweetsRepository: &TweetsMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)

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
	Tweets() ([]Tweet, error)
}

type TweetsMemoryRepository struct {
	tweets []Tweet
}

func (r *TweetsMemoryRepository) SaveTweet(tw Tweet) (int, error) {
	r.tweets = append(r.tweets, tw)

	return len(r.tweets), nil
}

func (r *TweetsMemoryRepository) Tweets() ([]Tweet, error) {
	return r.tweets, nil
}

type server struct {
	tweetsRepository TweetsRepository
}

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)

		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
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

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.tweetsRepository.Tweets()

	if err != nil {
		fmt.Println("Failed to get tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := tweetsList{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)

}
